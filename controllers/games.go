package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/IbisStats-Server/datasource"
	"github.com/lordmortis/IbisStats-Server/middleware"
	"github.com/pkg/errors"
)

func Games(router gin.IRoutes)  {
	router.GET("", listGames)
	router.POST("", createGame)
	router.GET("/:id", showGame)
	router.PUT("/:id", updateGame)
	router.DELETE("/:id", deleteGame)
}

func listGames(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	var models []datasource.Game
	if superAdmin {
		models, err = datasource.GamesAll(ctx)
		if err != nil {
			JSONInternalServerError(ctx, err)
			return
		}
	} else {
		JSONForbiddenResponse(ctx)
		return
	}

	JSONOk(ctx, models)
}


func createGame(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	model := datasource.Game{}

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

func showGame(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	model, err := datasource.GameWithIDString(ctx, ctx.Param("id"))
	if err != nil && err != datasource.ErrNoResults {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	if err == datasource.ErrNoResults {
		JSONNotFound(ctx)
	} else {
		JSONOk(ctx, model)
	}
}

func updateGame(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	model, err := datasource.GameWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	if model == nil {
		JSONNotFound(ctx)
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

func deleteGame(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	model, err := datasource.GameWithIDString(ctx, ctx.Param("id"))
	if err != nil && err != datasource.ErrNoResults {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
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