package db

import (
	"time"
)

type DailyAct struct {
	ActID         int       `xorm:"not null 'act_id'" json:"-"`
	UserID        int       `xorm:"not null 'user_id'" json:"-"`
	ActType       int       `xorm:"act_type" json:"act_type"`
	StartTime     time.Time `xorm:"-" json:"-"`
	StartTimeUnix int64     `xorm:"bigint 'start_time'" json:"start_time"`
	EndTime       time.Time `xorm:"-" json:"-"`
	EndTimeUnix   int64     `xorm:"'end_time'" json:"end_time"`
	Duration      string    `xorm:"char(32)"`
	Num           int
	Comment       string
}

// TableName return table name
func (act DailyAct) TableName() string {
	return "daily_act"
}

// CreateNewActivity create a new activity
func CreateNewActivity(act *DailyAct) error {
	_, err := x.Insert(act)
	return err
}
