package auth

import (
	cryptoRand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/lordmortis/IbisStats-Server/datasource"
	"github.com/pkg/errors"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID uuid.UUID
	Base64ID string
	Expiry time.Time
	model string
	modelID []byte
}

type redisSessionData struct {
	Model string
	ModelID []byte
	Secret []byte
}

var (
	expiryDuration time.Duration
)

func init() {
	expiryDuration = time.Minute * 15
	reader := cryptoRand.Reader
	seedBytes := []byte{0,0,0,0,0,0,0,0}
	num, err := reader.Read(seedBytes)
	if num < 8 || err != nil {
		panic("Unable to generate randomized seed!")
	}
	seed := binary.BigEndian.Uint64(seedBytes)
	rand.Seed(int64(seed))
}

func CreateNewSession(ctx *gin.Context, model string, modelID []byte) (*Session, error) {
	redisClient := ctx.MustGet("redisConnection").(*redis.Client)
	redisPrefix := ctx.MustGet("redisPrefix").(string)

	sessionID, _ := uuid.NewV4()
	base64ID := base64.StdEncoding.EncodeToString(sessionID.Bytes())
	secret := []byte{0,0,0,0,0,0,0,0}
	readLength, err := rand.Read(secret)
	if readLength != len(secret) {
		return nil, errors.New("unable to generate session secret")
	} else if err != nil {
		return nil, err
	}

	base64Secret := base64.StdEncoding.EncodeToString(secret)
	redisSession := redisSessionData{model, modelID, secret}
	encodedRedisSession, err := json.Marshal(redisSession)
	if err != nil {
		return nil, errors.Wrap(err, "Could not serialize session")
	}

	redisClient.Set(redisPrefix + "session:" + base64ID, encodedRedisSession, expiryDuration)
	maxAge := int(math.Round(expiryDuration.Seconds()))
	domain := strings.Split(ctx.Request.Host, ":")[0]
	ctx.SetCookie("sessionSecret", base64Secret, maxAge, "/", domain, false, true)

	session := Session{sessionID, base64ID, time.Now().Add(expiryDuration),  model, modelID}
	return &session, nil
}

func GetSession(ctx *gin.Context) (*Session, error)  {
	modelnameInf, ok := ctx.Get("AuthModelType")
	if !ok {
		return nil, errors.New("No Auth Data")
	}

	modelName, ok := modelnameInf.(string)
	if !ok {
		return nil, errors.New("Model name invalid")
	}

	idBytesInf, ok := ctx.Get("AuthModelID")
	if !ok {
		return nil, errors.New("No Auth Data")
	}

	idBytes, ok := idBytesInf.([]byte)
	if !ok {
		return nil, errors.New("No Auth Data")
	}

	sessionIDinf, ok := ctx.Get("SessionID")
	if !ok {
		return nil, errors.New("No Session ID Data")
	}

	sessionIDbase64String, ok := sessionIDinf.(string)
	if !ok {
		return nil, errors.New("No Session ID Data")
	}

	sessionIDBytes, err := base64.StdEncoding.DecodeString(sessionIDbase64String)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode session ID")
	}

	sessionID, err := uuid.FromBytes(sessionIDBytes)
	if !ok {
		return nil, errors.New("No Session ID Data")
	}

	session := Session{
		ID:       sessionID,
		Base64ID: sessionIDbase64String,
		Expiry:   time.Time{},
		model:    modelName,
		modelID:  idBytes,
	}

	return &session, nil
}

func (session *Session)IsSuperAdmin(ctx *gin.Context) (bool, error) {
	// todo: we need to map these onto proper datasource functions...
	modelUUID, err := uuid.FromBytes(session.modelID)
	if err != nil {
		return false, errors.New("Auth Data invalid - uuid can't be parsed")
	}

	user, err := datasource.UserWithUUID(ctx, modelUUID)
	if err != nil {
		return false, errors.New("Auth Data invalid - user model doesn't exist")
	}

	return user.SuperAdmin, nil
}

func (session *Session)IsModelOwner(ctx *gin.Context, modelInf interface{}) (bool, error) {
	idBase64 := base64.StdEncoding.EncodeToString(session.modelID)

	// todo: we need to map these onto proper datasource functions...
	if reflect.ValueOf(modelInf).IsNil() {
		return false, nil
	}

	model, ok := modelInf.(*datasource.User)
	if !ok {
		return false, nil
	}

	return model.ID == idBase64, nil
}

func (session *Session)Destroy(ctx *gin.Context) {
	redisClient := ctx.MustGet("redisConnection").(*redis.Client)
	redisPrefix := ctx.MustGet("redisPrefix").(string)

	redisClient.Del(fmt.Sprintf("%ssession:%s", redisPrefix, session.ID))
}
