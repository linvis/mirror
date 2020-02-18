package router

import (
	"mirror/api"
	"mirror/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterGroup(engine *gin.Engine, env string) {
	// home
	engine.GET("/", api.Home)
	engine.Use()

	user := engine.Group(env + "/user")
	{
		user.POST("/login", api.LoginIn)
		user.POST("/logout", api.Logout)
		user.GET("/info", api.GetUserInfo)
	}

	sub := engine.Group(env + "/record/submit")
	{
		sub.Use(middleware.BaseAuth())
		sub.POST("/sleep", api.NewSleepRecord)
	}

	query := engine.Group(env + "/record/query")
	{
		query.Use(middleware.BaseAuth())
		query.GET("/sleep", api.GetSleepRecord)
		query.GET("/leetcode", api.GetLeetcodeRecord)
		query.GET("/github", api.GetGithubRecord)
	}

	analysis := engine.Group(env + "/record/analysis")
	{
		analysis.Use(middleware.BaseAuth())
		analysis.GET("/sleep", api.GetSleepRecordAnalysis)
	}

	settings := engine.Group(env + "setting")
	{
		settings.Use(middleware.BaseAuth())
		settings.GET("/program", api.GetProgramSetting)
		settings.POST("/program", api.SetProgramSetting)
	}

	editor := engine.Group(env + "editor")
	{
		editor.Use(middleware.BaseAuth())
		editor.GET("/catalog", api.GetEditorCatalog)
		editor.POST("/catalog/update", api.UpdateEditorCatalog)
	}
}
