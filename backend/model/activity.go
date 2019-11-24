package model

import (
	"fmt"
	"mirror/database"
	"time"
)

type Activity struct {
	EventID   int       `gorm:"column:event_id"`
	UserID    int       `gorm:"column:user_id"`
	EventType int       `gorm:"column:event_type"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime   time.Time `gorm:"column:end_time"`
	Duration  string    `gorm:"column:duration"`
	Num       int       `gorm:column:"num"`
	Comment   string    `gorm:"column:comment"`
}

func (act Activity) TableName() string {
	return "activity"
}

func ActivityInsert(act Activity) {

	fmt.Println(act)

	db, err := database.GetDB()
	if err != nil {
		fmt.Println(err)
	}

	db.Create(&act)
}
