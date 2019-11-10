package main

import (
	"mirror/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	router.InitRouter(engine)

	engine.Static("/static", "./templates/static")
	engine.StaticFile("/favicon.ico", "./templates/favicon.ico")
	engine.LoadHTMLGlob("templates/*.html")

	engine.Run("localhost:9528")
}
