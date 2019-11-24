package router

import (
	"mirror/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, env string) {

	for _, url := range controller.URLHandle {
		info := url

		if info.Method == "POST" {
			engine.POST(env+info.URL, info.Handlers...)
			continue
		}

		if info.Method == "GET" {
			engine.GET(env+info.URL, info.Handlers...)
			continue
		}

	}
}
