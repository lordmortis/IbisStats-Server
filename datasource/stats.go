package datasource

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/IbisStats-Server/datamodels_raw"
	"github.com/volatiletech/sqlboiler/boil"
	"strings"
	"time"
)

type GameStat struct {
	game *Game
	ID string `json:"id"`
	Name string
	Type string
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	dbModel *datamodels_raw.GameStat
}

var validTypes = [...]string  {
	"score", "time",
}

func (viewModel *GameStat)Update(ctx *gin.Context) error {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return err
	}

	var rows int64
	if viewModel.updateDB() {
		err = viewModel.dbModel.Insert(ctx, dbCon, boil.Infer())
		if err == nil  {
			rows = 1
		} else {
			return err
		}
	} else {
		rows, err = viewModel.dbModel.Update(ctx, dbCon, boil.Infer())
	}

	if err := viewModel.dbModel.Reload(ctx, dbCon); err != nil {
		return err
	}

	viewModel.fromDB(viewModel.dbModel, viewModel.game)

	if rows == 0 {
		return ErrNoUpdate
	}

	return nil
}

func (viewModel *GameStat)Delete(ctx *gin.Context) error {
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

func (viewModel *GameStat)Validate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(viewModel.Name) == 0 {
		errorMap["name"] = []string{"must be present"}
	} else if len(viewModel.Name) < 4 {
		errorMap["name"] = []string{"must be at least 4 characters"}
	}

	viewModel.Type = strings.ToLower(viewModel.Type)
	validType := false
	for index := range validTypes {
		if validTypes[index] == viewModel.Type {
			validType = true
			break
		}
	}

	if !validType {
		errorMap["type"] = []string{"must be a valid type"}
	}

	return errorMap
}

func (viewModel *GameStat)fromDB(dbModel *datamodels_raw.GameStat, game *Game) {
	viewModel.game = game
	viewModel.dbModel = dbModel
	viewModel.ID = UUIDBase64FromString(dbModel.ID)
	viewModel.Name = dbModel.Name
	viewModel.Type = dbModel.Type

	if dbModel.CreatedAt.Valid {
		viewModel.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		viewModel.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}
}

func (viewModel *GameStat)updateDB() bool {
	newModel := false
	if viewModel.dbModel == nil {
		id, _ := uuid.NewV4()
		newModel = true
		viewModel.dbModel = &datamodels_raw.GameStat{
			ID: id.String(),
		}
	}

	viewModel.dbModel.Name = viewModel.Name
	if viewModel.game != nil {
		viewModel.dbModel.GameID = viewModel.game.dbModel.ID
	}
	viewModel.dbModel.Type = viewModel.Type

	return newModel
}

func convertGameStats(dbGameStats datamodels_raw.GameStatSlice, game *Game) []GameStat {
	gamesStats := make([]GameStat, len(dbGameStats))
	for index := range dbGameStats {
		gameStat := GameStat{}
		gameStat.fromDB(dbGameStats[index], game)
		gamesStats[index] = gameStat
	}
	return gamesStats
}