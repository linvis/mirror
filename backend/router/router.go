package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	engine.GET("/", homePage)
}

func homePage(c *gin.Context) {
	c.String(200, "hello world")
}
