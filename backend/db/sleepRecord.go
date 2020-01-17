package db

import (
	"errors"
	"fmt"
	"time"
)

CREATE TABLE sleep_record
(
  record_id INT(11)
  UNSIGNED ZEROFILL NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT
  (11) NOT NULL,
	record_date INT NOT NULL,
  start_time MEDIUMINT DEFAULT NULL,
  end_time MEDIUMINT DEFAULT NULL,
	duration MEDIUMINT NOT NULL
);


type SleepRecord struct {
	RecordID         int       `xorm:"unsigned zerofill not null 'record_id'" json:"-"`
	UserID        int       `xorm:"not null 'user_id'" json:"-"`
	EvtDate       int     `xorm:"not null 'record_date'" json:"record_date"`
	StartTime     time.Time `xorm:"-" json:"-"`
	StartTimeUnix int       `xorm:"int 'start_time'" json:"start_time"`
	EndTime       time.Time `xorm:"-" json:"-"`
	EndTimeUnix   int       `xorm:"'end_time'" json:"end_time"`
	Duration      int
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

func GetDailyEvt(user_id int, evt_id int, item_id int) ([]DailyEvt, error) {
	evts := []DailyEvt{}

	err := x.Where("user_id = ?", user_id).And("evt_type= ?", evt_id).And("evt_item = ?", item_id).Find(&evts)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no user daily evts")
	}

	return evts, nil
}
