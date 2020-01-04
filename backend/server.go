package main

import (
	"mirror/db"
	"mirror/router"
	"os"

	"github.com/gin-gonic/gin"
)

const DEV_URL = "dev-api/"

func main() {
	environment := ""
	dbPath := ""
	if len(os.Args) <= 2 {

	} else {
		if os.Args[1] == "dev" {
			environment = DEV_URL
		}

		dbPath = os.Args[2]
	}

	engine := gin.Default()

	router.InitRouter(engine, environment)
	db.InitDB(dbPath)

	engine.Static("/static", "./templates/static")
	engine.StaticFile("/favicon.ico", "./templates/favicon.ico")
	engine.LoadHTMLGlob("templates/*.html")

	engine.Run(":9528")

}
