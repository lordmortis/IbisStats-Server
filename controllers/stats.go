package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Stats(router gin.IRoutes)  {
	router.GET("/:game_id/stats", listStats)
	router.POST("/:game_id/stats", createStat)
	router.GET("/:game_id/stat/:id", showStat)
	router.PUT("/:game_id/stat/:id", updateStat)
	router.DELETE("/:game_id/stat/:id", deleteStat)
}

func listStats(ctx *gin.Context) {
	gameID := ctx.Param("game_id")

	ctx.JSON(http.StatusOK, gin.H{"Game ID:": gameID, "action": "list stats"})
}

func createStat(ctx *gin.Context) {
	gameID := ctx.Param("game_id")

	ctx.JSON(http.StatusOK, gin.H{"Game ID:": gameID, "action": "create stat"})
}

func showStat(ctx *gin.Context) {
	gameID := ctx.Param("game_id")
	statID := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{"Game ID": gameID, "Stat ID": statID, "action": "show stat"})
}

func updateStat(ctx *gin.Context) {
	gameID := ctx.Param("game_id")
	statID := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{"Game ID": gameID, "Stat ID": statID, "action": "update stat"})

}

func deleteStat(ctx *gin.Context) {
	gameID := ctx.Param("game_id")
	statID := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{"Game ID": gameID, "Stat ID": statID, "action": "delete stat"})
}