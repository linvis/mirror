package main

import (
	"mirror/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	router.InitRouter(engine)

	engine.Run("localhost:5000")
}
