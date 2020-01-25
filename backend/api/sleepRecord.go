package api

import (
	"encoding/json"
	"fmt"
	"mirror/db"
	"net/http"

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

	fmt.Println(rec)

	rec.UserID = 3
	err := db.NewSleepRecord(&rec)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

// {"act_type":"read","start_time":1575207586000,"end_time":1575207588000,"duration":"","num":1,"comment":"zzz"}

func buildSleepRecordMsg(records *[]db.SleepRecord) *map[string][]int64 {
	msg := map[string][]int64{
		"date":       []int64{},
		"duration":   []int64{},
		"start_time": []int64{},
		"end_time":   []int64{},
	}

	for _, rec := range *records {
		msg["date"] = append(msg["date"], rec.Date)
		msg["duration"] = append(msg["duration"], int64(rec.Duration))
		msg["start_time"] = append(msg["start_time"], int64(rec.StartTime))
		msg["end_time"] = append(msg["end_time"], int64(rec.EndTime))
	}

	return &msg
}

func GetSleepRecord(c *gin.Context) {
	// get last month

	msg, found := db.GetSleepRecordFromRedis(3)
	if found {
		fmt.Println("fetch sleep records from redis", msg)
		c.String(http.StatusOK, msg)
	} else {
		records, err := db.GetSleepRecord(3, 30)
		fmt.Println("fetch sleep records from db", records)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
			return
		}

		msg := buildSleepRecordMsg(records)

		b, _ := json.Marshal(gin.H{"code": 20000, "data": *msg})

		err = db.SetSleepRecordToRedis(3, string(b))
		fmt.Println(err)

		c.String(http.StatusOK, string(b))
	}
}

func init() {
	// RegisterURL("record/submit/sleep", "POST", NewSleepRecord)
	// RegisterURL("record/query/sleep/:time", "GET", getSleepRecords)
}
