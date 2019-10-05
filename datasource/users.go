package datasource

import (
	"database/sql"
	"encoding/base64"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/lordmortis/IbisStats-Server/datamodels_raw"
)

type User struct {
	ID string `json:"id"`
	Username string
	Email string
	SuperAdmin bool `json:"super_admin"`
	OldPassword string `json:"current_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

var (
	emailRegexp *regexp.Regexp
)

func init() {
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
}

func UsersWithUsername(ctx *gin.Context, username *string) (datamodels_raw.UserSlice, error){
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return nil, err
	}

	models, err := datamodels_raw.Users(qm.Where("username = ?", username)).All(ctx, dbCon)
	if err == sql.ErrNoRows {
		return nil, ErrNoResults
	}

	return models, err
}

func UserCreate(ctx *gin.Context, user User) (*User, error) {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return nil, err
	}

	dbModel := datamodels_raw.User{}
	user.ToDB(&dbModel)

	if err := dbModel.Insert(ctx, dbCon, boil.Infer()); err != nil {
		return nil, err
	}

	if err := dbModel.Reload(ctx, dbCon); err != nil {
		return nil, err
	}

	user = User{}
	user.FromDB(&dbModel)

	return &user, nil
}

func UsersAll(ctx *gin.Context) (datamodels_raw.UserSlice, error){
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	return datamodels_raw.Users().All(ctx, dbCon)
}

func UserWithID(ctx *gin.Context, stringID string) (*datamodels_raw.User, error){
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	userID := UUIDFromString(stringID)

	if userID == uuid.Nil {
		return nil, ErrUUIDParse
	}

	user, err := datamodels_raw.FindUser(ctx,dbCon, userID.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func UserUUID(user *datamodels_raw.User) uuid.UUID {
	return uuid.FromStringOrNil(user.ID)
}

func UserUUIDBase64(user *datamodels_raw.User) string {
	uuid := UserUUID(user)
	return base64.StdEncoding.EncodeToString(uuid.Bytes())
}

func UserSetPassword(user *datamodels_raw.User, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.EncryptedPassword = null.BytesFrom(hashedPassword)
	return nil
}

func UserValidatePassword(user *datamodels_raw.User, password string) bool {
	if err := bcrypt.CompareHashAndPassword(user.EncryptedPassword.Bytes, []byte(password)); err != nil {
		return false
	}
	return true
}

func (user *User)FromDB(dbModel *datamodels_raw.User) {
	user.ID = UserUUIDBase64(dbModel)
	user.Username = dbModel.Username
	user.Email = dbModel.Email
	user.SuperAdmin = dbModel.SuperAdmin.Bool

	if dbModel.CreatedAt.Valid {
		user.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		user.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}
}

func (user *User) ValidateCreate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(user.Username) == 0 {
		errorMap["username"] = []string{"must be present"}
	} else if len(user.Username) < 4 {
		errorMap["username"] = []string{"must be at least 4 characters"}
	}

	if len(user.Email) == 0 {
		errorMap["email"] = []string{"must be present"}
	} else if !emailRegexp.MatchString(user.Email) {
		errorMap["email"] = []string{"must be a valid email address"}
	}

	if len(user.NewPassword) == 0 {
		errorMap["new_password"] = []string{"required"}
		errorMap["password_confirmation"] = []string{"required"}
	} else if user.NewPassword != user.PasswordConfirmation {
		errorMap["new_password"] = []string{"must equal password_confirmation"}
		errorMap["password_confirmation"] = []string{"must equal new_password"}
	}

	return errorMap
}

func (user *User) ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(user.Username) != 0 && len(user.Username) < 4 {
		errorMap["username"] = []string{"must be at least 4 characters"}
	}

	if len(user.Email) > 0 && !emailRegexp.MatchString(user.Email) {
		errorMap["email"] = []string{"must be a valid email address"}
	}

	if len(user.NewPassword) > 0 && user.NewPassword != user.PasswordConfirmation {
		errorMap["new_password"] = []string{"must equal password_confirmation"}
		errorMap["password_confirmation"] = []string{"must equal new_password"}
	}

	return errorMap
}

func (user *User) ToDB(dbModel *datamodels_raw.User) {
	if len(dbModel.ID) == 0 {
		id, _ := uuid.NewV4()
		dbModel.ID = id.String()
	}

	if len(user.Email) > 0 {
		dbModel.Email = user.Email
	}

	if len(user.Username) > 0 {
		dbModel.Username = user.Username
	}

	if len(user.NewPassword) > 0 && len(user.PasswordConfirmation) > 0 && user.NewPassword == user.PasswordConfirmation {
		_ = UserSetPassword(dbModel, user.NewPassword)
	}
}