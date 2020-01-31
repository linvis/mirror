package api

import (
	"fmt"
	"mirror/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProgramSettings struct {
	db.LeetcodeSetting
	db.GithubSetting
}

func SetProgramSetting(c *gin.Context) {
	var s ProgramSettings

	if err := c.BindJSON(&s); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Setting"})
		return
	}

	fmt.Println(s)

	s.UserID1 = 3
	err := db.NewLeetcodeSetting(&s.LeetcodeSetting)
	if err != nil {
		fmt.Println(err)
	}

	s.UserID2 = 3
	db.NewGithubSetting(&s.GithubSetting)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

func GetProgramSetting(c *gin.Context) {
	var s ProgramSettings

	s.LeetcodeSetting, _ = db.GetLeetcodeSetting(3)

	s.GithubSetting, _ = db.GetGithubSetting(3)

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": s})
}
