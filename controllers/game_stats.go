package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/IbisStats-Server/auth"
	"github.com/lordmortis/IbisStats-Server/datasource"
	"github.com/pkg/errors"
)

func Stats(router gin.IRoutes)  {
	router.GET("/:game_id/stats", listStats)
	router.POST("/:game_id/stats", createStat)
	router.GET("/:game_id/stat/:id", showStat)
	router.PUT("/:game_id/stat/:id", updateStat)
	router.DELETE("/:game_id/stat/:id", deleteStat)
}

func listStats(ctx *gin.Context) {
	session, err := auth.GetSession(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	superAdmin, err := session.IsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	game, err := datasource.GameWithIDString(ctx, ctx.Param("game_id"))

	if game == nil {
		JSONNotFound(ctx)
		return
	} else if err != nil  {
		JSONInternalServerError(ctx, err)
		return
	}

	models, err := game.StatsAll(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	JSONOk(ctx, models)
}

func createStat(ctx *gin.Context) {
	session, err := auth.GetSession(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	superAdmin, err := session.IsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	game, err := datasource.GameWithIDString(ctx, ctx.Param("game_id"))

	if game == nil {
		JSONNotFound(ctx)
		return
	} else if err != nil  {
		JSONInternalServerError(ctx, err)
		return
	}

	model := game.CreateStat()

	if err := ctx.ShouldBindJSON(&model); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := model.Validate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	err = model.Update(ctx)
	if err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	JSONOk(ctx, model)
}

func showStat(ctx *gin.Context) {
	session, err := auth.GetSession(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	superAdmin, err := session.IsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	game, err := datasource.GameWithIDString(ctx, ctx.Param("game_id"))

	if game == nil {
		JSONNotFound(ctx)
		return
	} else if err != nil  {
		JSONInternalServerError(ctx, err)
		return
	}

	model, err := game.StatWithStringID(ctx, ctx.Param("id"))
	if err != nil && err != datasource.ErrNoResults {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if err == datasource.ErrNoResults {
		JSONNotFound(ctx)
	} else {
		JSONOk(ctx, model)
	}
}

func updateStat(ctx *gin.Context) {
	session, err := auth.GetSession(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	superAdmin, err := session.IsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	game, err := datasource.GameWithIDString(ctx, ctx.Param("game_id"))

	if game == nil {
		JSONNotFound(ctx)
		return
	} else if err != nil  {
		JSONInternalServerError(ctx, err)
		return
	}

	model, err := game.StatWithStringID(ctx, ctx.Param("id"))
	if err != nil && err != datasource.ErrNoResults {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if err := ctx.ShouldBindJSON(&model); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	modelErrors := model.Validate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	err = model.Update(ctx)
	if err != nil {
		if err == datasource.ErrNoUpdate {
			JSONBadRequest(ctx, gin.H{"general": [1]string{"no rows updated"}})
			return
		}
		JSONInternalServerError(ctx, err)
		_ = ctx.Error(err)
		return
	}

	JSONOk(ctx, model)
}

func deleteStat(ctx *gin.Context) {
	session, err := auth.GetSession(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	superAdmin, err := session.IsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	game, err := datasource.GameWithIDString(ctx, ctx.Param("game_id"))

	if game == nil {
		JSONNotFound(ctx)
		return
	} else if err != nil  {
		JSONInternalServerError(ctx, err)
		return
	}

	model, err := game.StatWithStringID(ctx, ctx.Param("id"))
	if err != nil && err != datasource.ErrNoResults {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if err == datasource.ErrNoResults || model == nil{
		JSONNotFound(ctx)
	} else {
		err = model.Delete(ctx)
		if err == datasource.ErrNotInDb || err == datasource.ErrNoResults {
			JSONNotFound(ctx)
		} else if err != nil {
			JSONInternalServerError(ctx, err)
		} else {
			JSONOkStatusResponse(ctx)
		}
	}
}