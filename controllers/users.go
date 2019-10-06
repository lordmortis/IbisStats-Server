package controllers

import (
	"github.com/lordmortis/IbisStats-Server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/lordmortis/IbisStats-Server/datasource"
)

func Users(router gin.IRoutes) {
	router.GET("", listUsers)
	router.POST("", createUsers)
	router.GET("/:id", showUsers)
	router.PUT("/:id", updateUsers)
	router.DELETE("/:id", deleteUsers)
}

func listUsers(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	var models []datasource.User

	if superAdmin {
		models, err = datasource.UsersAll(ctx)
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

func showUsers(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

 	model, err := datasource.UserWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

 	if superAdmin {
		if model == nil {
			JSONNotFound(ctx)
			return
		}
	} else {
		currentUser, err := middleware.AuthIsCurrentUser(ctx, model)
		if err != nil {
			JSONInternalServerError(ctx, err)
			return
		}
		if model == nil || !currentUser {
			JSONForbiddenResponse(ctx)
			return
		}
	}

	JSONOk(ctx, model)
}

func createUsers(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if !superAdmin {
		JSONForbiddenResponse(ctx)
		return
	}

	newJson := datasource.User{}

	if err := ctx.ShouldBindJSON(&newJson); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := newJson.ValidateCreate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	savedUser, err := datasource.UserCreate(ctx, newJson)
	if err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	JSONOk(ctx, savedUser)
}

func updateUsers(ctx *gin.Context) {
	superAdmin, err := middleware.AuthIsSuperAdmin(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	model, err := datasource.UserWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		if err == datasource.ErrUUIDParse {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if !superAdmin {
		currentUser, err := middleware.AuthIsCurrentUser(ctx, model)
		if err != nil {
			JSONInternalServerError(ctx, err)
			return
		}
		if model == nil || !currentUser {
			JSONForbiddenResponse(ctx)
			return
		}
	}

	if model == nil {
		JSONNotFound(ctx)
		return
	}

	updateModel := datasource.User{}
	if err := ctx.ShouldBindJSON(&updateModel); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	modelErrors := updateModel.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	updated, err := model.Update(ctx, updateModel)
	if err != nil {
		JSONInternalServerError(ctx, err)
		_ = ctx.Error(err)
		return
	}

	if updated {
		JSONOk(ctx, model)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"no rows updated"}})
	}
}

func deleteUsers(ctx *gin.Context) {
	model, err := datasource.UserWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if model == nil {
		JSONNotFound(ctx)
		return
	}

	deleted, err := model.Delete(ctx)

	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if deleted {
		JSONOkStatusResponse(ctx)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"unable to deleteUsers"}})
	}
}