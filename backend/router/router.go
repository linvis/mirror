package router

import (
	"mirror/api"

	"github.com/gin-gonic/gin"
)

func InitRouterGroup(engine *gin.Engine, env string) {
	// home
	engine.GET("/", api.Home)

	user := engine.Group(env + "/user")
	{
		user.POST("/login", api.LoginIn)
		user.POST("/logout", api.Logout)
		user.GET("/info", api.GetUserInfo)
	}

	sub := engine.Group(env + "/record/submit")
	{
		sub.POST("sleep", api.NewSleepRecord)
		sub.POST("leetcode", api.NewLeetcodeRecord)
	}

	query := engine.Group(env + "/record/query")
	{
		query.GET("sleep", api.GetSleepRecord)
		query.GET("leetcode", api.GetLeetcodeRecord)
	}

	analysis := engine.Group(env + "/record/analysis")
	{
		analysis.GET("sleep", api.GetSleepRecordAnalysis)
	}
}
