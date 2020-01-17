package controller

import (
	"encoding/json"
	"fmt"
	"mirror/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EvtItems struct {
	ID       int        `json:"value"`
	Name     string     `json:"label"`
	Children []EvtItems `json:"children, omitempty"`
}

func NewSleepRecord(c *gin.Context) {
	var rec db.SleepRecord

	if err := c.BindJSON(&rec); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	rec.UserID = 3
	db.NewSleepRecord(&rec)

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

// {"act_type":"read","start_time":1575207586000,"end_time":1575207588000,"duration":"","num":1,"comment":"zzz"}

func buildSleepDataMsg(evts *[]db.SleepRecord) string {
	date := []string{}
	duration := []int{}

	for _, e := range *evts {
		date = append(date, time.Unix(int64(e.Date), 0).Format("2006-1-2"))
		duration = append(duration, e.Duration)
	}

	b, _ := json.Marshal(gin.H{"code": 20000, "data": gin.H{"date": date, "duration": duration}})

	return string(b)
}

func getSleepRecords(c *gin.Context) {
	evts, err := db.GetSleepRecord(3)
	fmt.Println(evts)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	msg := buildSleepDataMsg(evts)

	c.String(http.StatusOK, msg)
}

func init() {
	RegisterURL("record/sleep/new", "POST", NewSleepRecord)
	RegisterURL("record/sleep/get", "GET", getSleepRecords)
}
