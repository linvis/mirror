package db

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

const (
	LastWeek = iota
	LastTwoWeek
	LastMonth
)

type SleepRecord struct {
	RecordID      int       `xorm:"unsigned zerofill not null 'record_id'" json:"-"`
	UserID        int       `xorm:"not null 'user_id'" json:"-"`
	Date          int64     `xorm:"bigint not null 'record_date'" json:"record_date"`
	StartTimeUnix time.Time `xorm:"-" json:"-"`
	StartTime     int       `xorm:"'start_time'" json:"start_time"`
	EndTimeUnix   time.Time `xorm:"-" json:"-"`
	EndTime       int       `xorm:"'end_time'" json:"end_time"`
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

func SetSleepRecordToRedis(user_id int, val interface{}) error {
	err := r.Set(strconv.Itoa(user_id), val, time.Hour*24).Err()

	return err
}

func GetSleepRecordFromRedis(user_id int) (string, bool) {
	found := false
	id := strconv.Itoa(user_id)

	val, err := r.Get(id).Result()

	if err == redis.Nil {
		// not found
	} else if err != nil {
		// error
		fmt.Println(err)
	} else {
		found = true
	}

	return val, found
}

func GetSleepRecord(user_id int, days int) (*[]SleepRecord, error) {
	rec := &[]SleepRecord{}

	err := x.Where("user_id = ? AND FROM_UNIXTIME(record_date/1000) >= DATE(NOW()) - INTERVAL ? DAY", user_id, days).Find(rec)
	// err := x.Where("user_id = ?", user_id).Find(rec)
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
