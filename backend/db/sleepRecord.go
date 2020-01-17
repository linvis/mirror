package db

import (
	"errors"
	"fmt"
	"time"
)

type SleepRecord struct {
	RecordID      int       `xorm:"unsigned zerofill not null 'record_id'" json:"-"`
	UserID        int       `xorm:"not null 'user_id'" json:"-"`
	Date          int       `xorm:"not null 'record_date'" json:"record_date"`
	StartTime     time.Time `xorm:"-" json:"-"`
	StartTimeUnix int       `xorm:"'start_time'" json:"start_time"`
	EndTime       time.Time `xorm:"-" json:"-"`
	EndTimeUnix   int       `xorm:"'end_time'" json:"end_time"`
	Duration      int       `xorm:"not null"`
}

// TableName return table name
func (r SleepRecord) TableName() string {
	return "sleep_record"
}

func NewSleepRecord(rec *SleepRecord) error {
	_, err := x.Insert(rec)
	return err
}

func GetSleepRecord(user_id int) (*[]SleepRecord, error) {
	rec := &[]SleepRecord{}

	err := x.Where("user_id = ?", user_id).Find(rec)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no user daily evts")
	}

	return rec, nil
}

func DeleteSleepRecord(id int) error {
	rec := SleepRecord{}

	_, err := x.Where("record_id = ?", id).Delete(&rec)
	if err != nil {
		return errors.New("delete rec fail")
	}

	return nil
}
