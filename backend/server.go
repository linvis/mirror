package main

import (
	"mirror/db"
	"mirror/middleware"
	"mirror/router"
	_ "mirror/spider"
	"os"

	"github.com/gin-gonic/gin"
)

const DEV_URL = "dev-api/"

func main() {
	environment := ""
	dbPath := ""
	redisPath := ""
	if len(os.Args) <= 3 {

	} else {
		if os.Args[1] == "dev" {
			environment = DEV_URL
		}

		dbPath = os.Args[2]
		redisPath = os.Args[3]
	}

	gin.SetMode(gin.DebugMode)

	middleware.InitLogger(environment)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	router.InitRouterGroup(engine, environment)

	db.InitDB(dbPath)
	db.InitRedis(redisPath)
	db.InitMongoDB()

	engine.Static("/static", "./templates/static")
	engine.StaticFile("/favicon.ico", "./templates/favicon.ico")
	engine.LoadHTMLGlob("templates/*.html")

	engine.Run(":9528")

}
