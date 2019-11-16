package datasource

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/IbisStats-Server/datamodels_raw"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type Game struct {
	ID string `json:"id"`
	Name string
	Secret string
	binarySecret []byte
	RegenerateSecret bool `json:"regenerate_secret,omitempty"`
	dbModel *datamodels_raw.Game
}

func GamesAll(ctx *gin.Context)  ([]Game, error) {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return nil, err
	}

	dbModels, err := datamodels_raw.Games().All(ctx, dbCon)
	if err == sql.ErrNoRows {
		return nil, ErrNoResults
	}

	return convertGames(dbModels), nil
}

func GamesWithName(ctx *gin.Context, name *string) ([]Game, error) {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return nil, err
	}

	dbModels, err := datamodels_raw.Games(qm.Where("name = ?", name)).All(ctx, dbCon)
	if err == sql.ErrNoRows {
		return nil, ErrNoResults
	}

	return convertGames(dbModels), nil
}

func GameWithID(ctx *gin.Context, id uuid.UUID) (*Game, error) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	dbModel, err := datamodels_raw.FindGame(ctx,dbCon, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := Game{}
	model.fromDB(dbModel)

	return &model, nil
}

func GameWithIDString(ctx *gin.Context, stringID string) (*Game, error) {
	userID := UUIDFromString(stringID)

	if userID == uuid.Nil {
		return nil, ErrUUIDParse
	}

	return GameWithID(ctx, userID)
}

func (game *Game)Update(ctx *gin.Context) error {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return err
	}

	var rows int64
	if game.updateDB() {
		err = game.dbModel.Insert(ctx, dbCon, boil.Infer())
		if err == nil  {
			rows = 1
		}
	} else {
		rows, err = game.dbModel.Update(ctx, dbCon, boil.Infer())
	}

	if err := game.dbModel.Reload(ctx, dbCon); err != nil {
		return err
	}

	game.fromDB(game.dbModel)

	if rows == 0 {
		return ErrNoUpdate
	}

	return nil
}

func (game *Game)Delete(ctx *gin.Context) error {
	if game.dbModel == nil  {
		return ErrNotInDb
	}

	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return err
	}

	rows, err := game.dbModel.Delete(ctx, dbCon)
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNoResults
	}

	return nil
}

func (game *Game)Validate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(game.Name) == 0 {
		errorMap["name"] = []string{"must be present"}
	} else if len(game.Name) < 4 {
		errorMap["name"] = []string{"must be at least 4 characters"}
	}

	return errorMap
}

func (game *Game)fromDB(dbModel *datamodels_raw.Game) {
	game.ID = UUIDBase64FromString(dbModel.ID)
	game.Name = dbModel.Name
	game.Secret = base64.StdEncoding.EncodeToString(dbModel.Secret)
	game.binarySecret = dbModel.Secret
	game.dbModel = dbModel
}

func (game *Game)updateDB() bool {
	newModel := false
	if game.dbModel == nil {
		id, _ := uuid.NewV4()
		newModel = true
		game.dbModel = &datamodels_raw.Game{
			ID: id.String(),
		}

		game.RegenerateSecret = true
	}

	game.dbModel.Name = game.Name
	if game.RegenerateSecret  {
		game.generateSecret()
		game.dbModel.Secret = game.binarySecret
	}

	return newModel
}

func (game *Game)generateSecret()  {
	game.binarySecret = make([]byte, 32)
	_, err := rand.Read(game.binarySecret)
	if err != nil {
		panic("could not generate secret!?")
	}
}

func convertGames(dbGames datamodels_raw.GameSlice) []Game {
	games := make([]Game, len(dbGames))
	for index := range dbGames {
		game := Game{}
		game.fromDB(dbGames[index])
		games[index] = game
	}
	return games
}