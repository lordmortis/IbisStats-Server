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
	"time"
)

type Game struct {
	ID string `json:"id"`
	Name string
	Secret string
	binarySecret []byte
	RegenerateSecret bool `json:"regenerate_secret,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
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

func (viewModel *Game)StatsAll(ctx *gin.Context) ([]GameStat, error) {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return nil, err
	}

	dbModels, err := datamodels_raw.GameStats(qm.Where("game_id = ?", viewModel.dbModel.ID)).All(ctx, dbCon)
	if err != nil {
		return nil, err
	}

	return convertGameStats(dbModels, viewModel), nil
}

func (viewModel *Game)StatWithID(ctx *gin.Context, id uuid.UUID) (*GameStat, error) {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return nil, err
	}

	dbModel, err := datamodels_raw.GameStats(qm.Where("game_id = ? AND id = ?", viewModel.dbModel.ID, id)).One(ctx, dbCon)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := GameStat{}
	model.fromDB(dbModel, viewModel)

	return &model, nil
}

func (viewModel *Game)StatWithStringID(ctx *gin.Context, stringID string) (*GameStat, error) {
	modelUUID := UUIDFromString(stringID)

	if modelUUID == uuid.Nil {
		return nil, ErrUUIDParse
	}

	return viewModel.StatWithID(ctx, modelUUID)
}

func (viewModel *Game)CreateStat() GameStat {
	return GameStat{game: viewModel}
}

func (viewModel *Game)Update(ctx *gin.Context) error {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return err
	}

	var rows int64
	if viewModel.updateDB() {
		err = viewModel.dbModel.Insert(ctx, dbCon, boil.Infer())
		if err == nil  {
			rows = 1
		}
	} else {
		rows, err = viewModel.dbModel.Update(ctx, dbCon, boil.Infer())
	}

	if err := viewModel.dbModel.Reload(ctx, dbCon); err != nil {
		return err
	}

	viewModel.fromDB(viewModel.dbModel)

	if rows == 0 {
		return ErrNoUpdate
	}

	return nil
}

func (viewModel *Game)Delete(ctx *gin.Context) error {
	if viewModel.dbModel == nil  {
		return ErrNotInDb
	}

	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return err
	}

	rows, err := viewModel.dbModel.Delete(ctx, dbCon)
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNoResults
	}

	return nil
}

func (viewModel *Game)Validate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(viewModel.Name) == 0 {
		errorMap["name"] = []string{"must be present"}
	} else if len(viewModel.Name) < 4 {
		errorMap["name"] = []string{"must be at least 4 characters"}
	}

	return errorMap
}

func (viewModel *Game)fromDB(dbModel *datamodels_raw.Game) {
	viewModel.ID = UUIDBase64FromString(dbModel.ID)
	viewModel.Name = dbModel.Name
	viewModel.Secret = base64.StdEncoding.EncodeToString(dbModel.Secret)
	viewModel.binarySecret = dbModel.Secret
	viewModel.dbModel = dbModel

	if dbModel.CreatedAt.Valid {
		viewModel.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		viewModel.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}
}

func (viewModel *Game)updateDB() bool {
	newModel := false
	if viewModel.dbModel == nil {
		id, _ := uuid.NewV4()
		newModel = true
		viewModel.dbModel = &datamodels_raw.Game{
			ID: id.String(),
		}

		viewModel.RegenerateSecret = true
	}

	viewModel.dbModel.Name = viewModel.Name
	if viewModel.RegenerateSecret  {
		viewModel.generateSecret()
		viewModel.dbModel.Secret = viewModel.binarySecret
	}

	return newModel
}

func (viewModel *Game)generateSecret()  {
	viewModel.binarySecret = make([]byte, 32)
	_, err := rand.Read(viewModel.binarySecret)
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