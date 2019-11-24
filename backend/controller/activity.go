package controller

import (
	"fmt"
	"mirror/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ActivityInfo struct {
	Type      string `json:"type"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Duration  string `json:"duration"`
	Num       int    `json:"num"`
	Comment   string `json:"comment"`
}

func newActivity(c *gin.Context) {

	// body, _ := ioutil.ReadAll(c.Request.Body)

	var activity ActivityInfo

	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	fmt.Println(activity)
	model.ActivityInsert(model.Activity{
		UserID:    3,
		EventType: 1,
		StartTime: time.Unix(activity.StartTime/1000, 0),
		EndTime:   time.Unix(activity.EndTime/1000, 0),
		Duration:  activity.Duration,
		Num:       activity.Num,
		Comment:   activity.Comment,
	})

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

func init() {
	RegisterURL("activity/new", "POST", newActivity)
}
