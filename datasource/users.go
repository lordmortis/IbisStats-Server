package datasource

import (
	"database/sql"
	"encoding/base64"
	"github.com/volatiletech/null"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/lordmortis/IbisStats-Server/datamodels_raw"
)

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