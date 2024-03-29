package main

import (
	"flag"
	"fmt"
	"github.com/lordmortis/IbisStats-Server/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"runtime"

	"github.com/lordmortis/IbisStats-Server/config"
	"github.com/lordmortis/IbisStats-Server/controllers"
	"github.com/lordmortis/IbisStats-Server/datasource"
	"github.com/lordmortis/IbisStats-Server/middleware"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configFile := flag.String("config", "config.json", "JSON Config file")

	flag.Parse()

	conf, err := config.Load(configFile)
	if err != nil {
		fmt.Println("Unable to parse Config file")
		fmt.Println(err)
		return
	}

	err = datasource.PerformMigrations(conf.Database, conf.Development)
	if err != nil {
		fmt.Println("Unable to perform/check migrations")
		fmt.Println(err)
		return
	}

	var dbMiddleware gin.HandlerFunc
	dbMiddleware, err = middleware.Database(conf.Database)
	if err != nil {
		fmt.Println("Unable to setup database connection:")
		fmt.Println(err)
		return
	}

	var redisMiddleware gin.HandlerFunc
	redisMiddleware, err = middleware.Redis(conf.Redis)
	if err != nil {
		fmt.Println("Unable to connect to redis:")
		fmt.Println(err)
		return
	}

	authMiddleware := auth.GinMiddleware()

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = conf.Server.AllowedOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))
	router.Use(dbMiddleware)
	router.Use(redisMiddleware)

	loginGroup := router.Group("/1/login")
	controllers.Login(loginGroup)

	sessionKeepalive := router.Group("/1/auth")
	sessionKeepalive.Use(authMiddleware)
	controllers.Auth(sessionKeepalive)

	userGroup := router.Group("/1/users")
	userGroup.Use(authMiddleware)
	controllers.Users(userGroup)

	gamesGroup := router.Group("/1/games")
	gamesGroup.Use(authMiddleware)
	controllers.Games(gamesGroup)

	gameGroup := router.Group("/1/game")
	gameGroup.Use(authMiddleware)
	controllers.Games(gameGroup)
	controllers.Stats(gameGroup)

	err = router.Run(conf.Server.String())
	if err != nil {
		fmt.Println("Unable to start server")
		fmt.Println(err)
		return
	}
}