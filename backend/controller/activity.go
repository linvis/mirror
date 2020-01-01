package controller

import (
	"fmt"
	"mirror/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// {"act_type":"read","start_time":1575207586000,"end_time":1575207588000,"duration":"","num":1,"comment":"zzz"}
func newActivity(c *gin.Context) {

	// body, _ := ioutil.ReadAll(c.Request.Body)
	var activity db.DailyAct

	if err := c.BindJSON(&activity); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	fmt.Println(activity)
	db.CreateNewActivity(&db.DailyAct{
		UserID:  3,
		ActType: activity.ActType,
		// StartTime: time.Unix(activity.StartTimeUnix/1000, 0),
		// EndTime:   time.Unix(activity.EndTimeUnix/1000, 0),
		StartTimeUnix: activity.StartTimeUnix / 1000,
		EndTimeUnix:   activity.EndTimeUnix / 1000,
		Duration:      activity.Duration,
		Num:           activity.Num,
		Comment:       activity.Comment,
	})

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

func getActType(c *gin.Context) {
	types, err := db.GetUserActType(3)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "internal error"})
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": types})
}

func init() {
	RegisterURL("activity/new", "POST", newActivity)
	RegisterURL("activity/type", "GET", getActType)
}
