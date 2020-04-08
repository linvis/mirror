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

type avgSleep struct {
	duration int
	start    int
	end      int
	count    int
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

		var avgWeek, avgTwoWeek, avgMonth avgSleep

		now := time.Now()

		fmt.Println(now.AddDate(0, 0, -7).Local().Unix())
		for i := 0; i < len(msg.Data.Date); i++ {

			if msg.Data.Date[i] >= now.AddDate(0, 0, -7).Local().Unix()*1000 {
				fmt.Println(msg.Data.Date[i])
				avgWeek.duration += msg.Data.Duration[i]
				avgWeek.start += msg.Data.StartTime[i]
				avgWeek.end += msg.Data.EndTime[i]
				avgWeek.count++
			} else if msg.Data.Date[i] >= now.AddDate(0, 0, -14).Local().Unix()*1000 {
				avgTwoWeek.duration += msg.Data.Duration[i]
				avgTwoWeek.start += msg.Data.StartTime[i]
				avgTwoWeek.end += msg.Data.EndTime[i]
				avgTwoWeek.count++
			} else {
				avgMonth.duration += msg.Data.Duration[i]
				avgMonth.start += msg.Data.StartTime[i]
				avgMonth.end += msg.Data.EndTime[i]
				avgMonth.count++
			}
		}

		if avgMonth.count <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "no data"})
			return
		} else {
			avgMonth.duration /= avgMonth.count
			avgMonth.start /= avgMonth.count
			avgMonth.end /= avgMonth.count
		}

		if avgTwoWeek.count > 0 {
			avgTwoWeek.duration /= avgTwoWeek.count
			avgTwoWeek.start /= avgTwoWeek.count
			avgTwoWeek.end /= avgTwoWeek.count
		}

		if avgWeek.count > 0 {
			avgWeek.duration /= avgWeek.count
			avgWeek.start /= avgWeek.count
			avgWeek.end /= avgWeek.count
		}

		c.JSON(http.StatusOK, gin.H{"code": 20000, "data": map[string][]int{
			"duration":   []int{avgWeek.duration, avgTwoWeek.duration, avgMonth.duration},
			"start_time": []int{avgWeek.start, avgTwoWeek.start, avgMonth.start},
			"end_time":   []int{avgWeek.end, avgTwoWeek.end, avgMonth.end},
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
