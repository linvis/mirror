package router

import (
	"mirror/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	engine.GET("/", controller.HomePage)
}
