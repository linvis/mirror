package db

import (
	"errors"
	"fmt"
	"time"
)

type DailyEvt struct {
	EvtID         int       `xorm:"not null 'evt_id'" json:"-"`
	UserID        int       `xorm:"not null 'user_id'" json:"-"`
	EvtType       int       `xorm:"not null evt_type" json:"evt_type"`
	EvtItem       int       `xorm:"not null evt_item" json:"evt_item"`
	EvtDate       int64     `xorm:"bigint not null 'evt_date'" json:"evt_date"`
	StartTime     time.Time `xorm:"-" json:"-"`
	StartTimeUnix int       `xorm:"int 'start_time'" json:"start_time"`
	EndTime       time.Time `xorm:"-" json:"-"`
	EndTimeUnix   int       `xorm:"'end_time'" json:"end_time"`
	Duration      int
	Num           int
	Comment       string
}

// TableName return table name
func (evt DailyEvt) TableName() string {
	return "daily_evt"
}

// CreateNewActivity create a new event
func CreateDailyEvt(evt *DailyEvt) error {
	_, err := x.Insert(evt)
	return err
}

func GetAllDailyEvt(user_id int) ([]DailyEvt, error) {
	evts := []DailyEvt{}

	err := x.Where("user_id = ?", user_id).Find(&evts)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no user daily evts")
	}

	return evts, nil
}

func DelDailyEvt(evt_id int) error {
	evt := DailyEvt{}

	_, err := x.Where("evt_id = ?", evt_id).Delete(&evt)
	if err != nil {
		return errors.New("delete evt fail")
	}

	return nil
}
