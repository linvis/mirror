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

func buildSleepDataMsg(records *[]db.SleepRecord) string {
	msg := [][]int64{}

	for _, rec := range *records {
		msg = append(msg, []int64{rec.Date, (int64)(rec.Duration)})
	}

	b, _ := json.Marshal(gin.H{"code": 20000, "data": msg})

	return string(b)
}

func GetSleepRecord(c *gin.Context) {
	// get last month

	msg, found := db.GetSleepRecordFromRedis(3)
	if found {
		fmt.Println("fetch sleep records from redis", msg)
		c.JSON(http.StatusOK, gin.H{"code": 20000, "data": msg})
	} else {
		msg := [][]int64{}

		records, err := db.GetSleepRecord(3, 30)
		fmt.Println("fetch sleep records from db", records)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
			return
		}

		for _, rec := range *records {
			msg = append(msg, []int64{rec.Date, (int64)(rec.Duration)})
		}

		b, _ := json.Marshal(msg)

		err = db.SetSleepRecordToRedis(3, string(b))
		fmt.Println(err)

		c.JSON(http.StatusOK, gin.H{"code": 20000, "data": msg})
	}
}

func init() {
	// RegisterURL("record/submit/sleep", "POST", NewSleepRecord)
	// RegisterURL("record/query/sleep/:time", "GET", getSleepRecords)
}
