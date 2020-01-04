package controller

import (
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

func formatEvtItems(types []db.DailyEvtType, items []db.DailyEvtItem) []EvtItems {
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

func getEvtType(c *gin.Context) {
	var err error

	types, err := db.GetDailyEvtType(3)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "internal error"})
		return
	}

	items, err := db.GetDailyEvtItem(3)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "internal error"})
		return
	}
	fmt.Println(items)

	evtItems := formatEvtItems(types, items)
	fmt.Println(gin.H{"data": evtItems})

	// c.JSON(http.StatusOK, gin.H{"code": 20000, "data": evtItems})
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": evtItems})
}

func init() {
	RegisterURL("evt/new", "POST", newEvt)
	RegisterURL("evt/type", "GET", getEvtType)
}
