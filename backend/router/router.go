package router

import (
	"mirror/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	controller.InitHomeRouter(engine)
	controller.InitUserRouter(engine)
}
