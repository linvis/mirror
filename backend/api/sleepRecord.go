package api

import (
	"encoding/json"
	"fmt"
	"mirror/db"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type EvtItems struct {
	ID       int        `json:"value"`
	Name     string     `json:"label"`
	Children []EvtItems `json:"children, omitempty"`
}

func NewSleepRecord(c *gin.Context) {
	var rec db.SleepRecord

	id := GetUserID(c)

	if err := c.BindJSON(&rec); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	fmt.Println(rec)

	rec.UserID = id
	err := db.NewSleepRecord(&rec)
	if err != nil {
		log.Error(err)
	}

	// delete redis record
	db.DeleteSleepRecordFromRedis(id)

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

// {"act_type":"read","start_time":1575207586000,"end_time":1575207588000,"duration":"","num":1,"comment":"zzz"}

type SortBySleepRecord []db.SleepRecord

func (a SortBySleepRecord) Len() int           { return len(a) }
func (a SortBySleepRecord) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBySleepRecord) Less(i, j int) bool { return a[i].Date < a[j].Date }

func buildSleepRecordMsg(records *[]db.SleepRecord) *map[string][]int64 {
	msg := map[string][]int64{
		"date":       []int64{},
		"duration":   []int64{},
		"start_time": []int64{},
		"end_time":   []int64{},
	}

	sort.Sort(SortBySleepRecord(*records))

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
	id := GetUserID(c)

	msg, found := db.GetSleepRecordFromRedis(id)
	if found {
		log.Debug("fetch sleep records from redis", msg)
		c.String(http.StatusOK, msg)
	} else {
		records, err := db.GetSleepRecord(id, 30)
		log.Debug("fetch sleep records from db", records)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
			return
		}

		msg := buildSleepRecordMsg(records)

		b, _ := json.Marshal(gin.H{"code": 20000, "data": *msg})

		err = db.SetSleepRecordToRedis(id, string(b))

		c.String(http.StatusOK, string(b))
	}
}

type SleepRecordMsg struct {
	Code int
	Data struct {
		Date      []int64
		Duration  []int
		StartTime []int `json:"start_time"`
		EndTime   []int `json:"end_time"`
	}
}

func GetSleepRecordAnalysis(c *gin.Context) {
	id := GetUserID(c)
	timeout := 0
RETRY:
	msgCache, found := db.GetSleepRecordFromRedis(id)
	if found {
		var msg SleepRecordMsg

		err := json.Unmarshal([]byte(msgCache), &msg)
		if err != nil {
			fmt.Println(err)
		}

		avgDuration := make([]int, 3)
		avgStartTime := make([]int, 3)
		avgEndTime := make([]int, 3)
		count := make([]int, 3)

		for i := 0; i < len(msg.Data.Date); i++ {
			if i < 7 {
				avgDuration[0] += msg.Data.Duration[i]
				avgStartTime[0] += msg.Data.StartTime[i]
				avgEndTime[0] += msg.Data.EndTime[i]
				count[0]++
			} else if i < 14 {
				avgDuration[1] += msg.Data.Duration[i]
				avgStartTime[1] += msg.Data.StartTime[i]
				avgEndTime[1] += msg.Data.EndTime[i]
				count[1]++
			} else {
				avgDuration[2] += msg.Data.Duration[i]
				avgStartTime[2] += msg.Data.StartTime[i]
				avgEndTime[2] += msg.Data.EndTime[i]
				count[2]++
			}
		}

		if count[0] <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "no data"})
			return
		} else {
			avgDuration[0] /= count[0]
			avgStartTime[0] /= count[0]
			avgEndTime[0] /= count[0]
		}
		// last week
		if count[1] > 0 {
			avgDuration[1] /= count[1]
			avgStartTime[1] /= count[1]
			avgEndTime[1] /= count[1]
		}
		// last month
		if count[2] > 0 {
			avgDuration[2] = avgDuration[1]
			avgStartTime[2] = avgStartTime[1]
			avgEndTime[2] = avgEndTime[1]
		}

		c.JSON(http.StatusOK, gin.H{"code": 20000, "data": map[string][]int{
			"duration":   avgDuration,
			"start_time": avgStartTime,
			"end_time":   avgEndTime,
		}})
		return
	} else {
		if timeout > 1000 {
			goto END
		}
		time.Sleep(50 * time.Millisecond)
		timeout += 50
		goto RETRY
	}
END:
	c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "data not found"})
}

func init() {
	// RegisterURL("record/submit/sleep", "POST", NewSleepRecord)
	// RegisterURL("record/query/sleep/:time", "GET", getSleepRecords)
}
