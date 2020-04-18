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

	engine.POST("/user/login", api.LoginIn)
	engine.POST("/user/logout", api.Logout)

	user := engine.Group(env + "/user")
	{
		user.Use(middleware.BaseAuth())
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
		editor.POST("/document/new", api.NewDocument)
		editor.POST("/document/delete/:docid", api.DeleteDocument)
		editor.GET("/document/queryall", api.GetAllDocument)
		editor.GET("/document/query/:docid", api.GetDocumentByID)
	}
}
