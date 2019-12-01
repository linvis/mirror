package main

import (
	"mirror/router"
	"os"

	"github.com/gin-gonic/gin"
)

const DEV_URL = "dev-api/"

func main() {
	environment := ""
	if len(os.Args) <= 1 {

	} else {
		if os.Args[1] == "dev" {
			environment = DEV_URL
		}
	}

	engine := gin.Default()

	router.InitRouter(engine, environment)

	engine.Static("/static", "./templates/static")
	engine.StaticFile("/favicon.ico", "./templates/favicon.ico")
	engine.LoadHTMLGlob("templates/*.html")

	engine.Run("localhost:9528")

}
