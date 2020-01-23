package router

import (
	"mirror/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, env string) {

	for _, url := range api.URLHandle {
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

func InitRouterGroup(engine *gin.Engine) {
	// home
	engine.GET("/", api.Home)

	sub := engine.Group("/dev-api/record/submit")
	{
		sub.POST("sleep", api.NewSleepRecord)
	}

	query := engine.Group("/dev-api/record/query")
	{
		query.GET("sleep", api.GetSleepRecord)
	}
}
