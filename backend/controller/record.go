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

// {"act_type":"read","start_time":1575207586000,"end_time":1575207588000,"duration":"","num":1,"comment":"zzz"}
func newEvt(c *gin.Context) {

	// body, _ := ioutil.ReadAll(c.Request.Body)
	var evt db.DailyEvt

	if err := c.BindJSON(&evt); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	fmt.Println(evt)
	db.CreateDailyEvt(&db.DailyEvt{
		UserID:  3,
		EvtType: evt.EvtType,
		EvtItem: evt.EvtItem,
		// StartTime: time.Unix(activity.StartTimeUnix/1000, 0),
		// EndTime:   time.Unix(activity.EndTimeUnix/1000, 0),
		EvtDate:       evt.EvtDate,
		StartTimeUnix: evt.StartTimeUnix,
		EndTimeUnix:   evt.EndTimeUnix,
		Duration:      evt.Duration,
		Num:           evt.Num,
		Comment:       evt.Comment,
	})

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

func formatEvtItemMsg(types []db.DailyEvtType, items []db.DailyEvtItem) []EvtItems {
	evtItems := []EvtItems{}

	for _, t := range types {
		evtItems = append(evtItems, EvtItems{t.EvtType, t.EvtName, nil})
	}

	for i, j := 0, 0; i < len(items) && j < len(evtItems); {
		if items[i].EvtType == evtItems[j].ID {
			if evtItems[j].Children == nil {
				evtItems[j].Children = []EvtItems{}
			}

			evtItems[j].Children = append(evtItems[j].Children, EvtItems{items[i].EvtItem, items[i].EvtItemName, nil})
			i++
		} else {
			j++
		}
	}
	// for i := 0; i < len(evtItems); i++ {
	// 	if evtItems[i].childrenArray != nil {
	// 		b, _ := json.Marshal(evtItems[i].childrenArray)
	// 		evtItems[i].Children = string(b)
	// 	}
	// }

	// b, _ := json.Marshal(evtItems)
	// fmt.Println(string(b))

	return evtItems
}

func buildDailyEvtItemMsg(user_id int) (string, error) {

	b, _ := json.Marshal(gin.H{"code": 60204, "message": "internal error"})
	msg := string(b)

	types, err := db.GetDailyEvtType(3)

	if err != nil {
		return msg, err
	}

	items, err := db.GetDailyEvtItem(3)
	if err != nil {
		return msg, err
	}

	evtItems := formatEvtItemMsg(types, items)

	b, _ = json.Marshal(gin.H{"code": 20000, "data": evtItems})

	return string(b), nil
}

func getEvtType(c *gin.Context) {
	msg, found := db.GetDailyEvtItemFromRedis(3)
	if found {
		fmt.Println("get from redis")
		c.String(http.StatusOK, msg)
	} else {
		// query from database

		msg, err := buildDailyEvtItemMsg(3)

		if err == nil {
			// put it in redis
			err = db.SetDailyEvtItemToRedis(3, msg)
			fmt.Println(err)
		}

		c.String(http.StatusOK, msg)
	}
}

func buildSleepDataMsg(evts []db.DailyEvt) string {
	date := []string{}
	duration := []int{}

	for _, e := range evts {
		date = append(date, time.Unix(e.EvtDate, 0).Format("2006-1-2"))
		duration = append(duration, e.Duration)
	}

	b, _ := json.Marshal(gin.H{"code": 20000, "data": gin.H{"date": date, "duration": duration}})

	return string(b)
}

func getSleepData(c *gin.Context) {
	evts, err := db.GetDailyEvt(3, 1, 0)
	fmt.Println(evts)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	msg := buildSleepDataMsg(evts)

	c.String(http.StatusOK, msg)
}

func init() {
	RegisterURL("evt/new", "POST", newEvt)
	RegisterURL("evt/type", "GET", getEvtType)
	RegisterURL("evt/sleep", "GET", getSleepData)
}
