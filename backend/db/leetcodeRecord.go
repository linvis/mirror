package db

import (
	"errors"
	"fmt"
)

type LeetcodeRecord struct {
	RecordID int   `xorm:"unsigned zerofill not null 'record_id'" json:"-"`
	UserID   int   `xorm:"not null 'user_id'" json:"-"`
	Date     int64 `xorm:"bigint not null 'record_date'" json:"record_date"`
	Num      int
}

func (r LeetcodeRecord) TableName() string {
	return "leetcode_records"
}

func NewLeetcodeRecord(rec *LeetcodeRecord) error {
	oldRec := LeetcodeRecord{}

	has, err := x.Where("user_id = ? AND record_date = ?", rec.UserID, rec.Date).Get(&oldRec)
	if err != nil {
		return err
	}

	if has {
		x.Update(rec, &LeetcodeRecord{RecordID: oldRec.RecordID})
	} else {
		_, err = x.Insert(rec)
	}

	return err
}

func GetLeetcodeRecord(user_id int, days int) (*[]LeetcodeRecord, error) {
	rec := &[]LeetcodeRecord{}

	err := x.Where("user_id = ? AND FROM_UNIXTIME(record_date/1000) >= DATE(NOW()) - INTERVAL ? DAY", user_id, days).Find(rec)
	// err := x.Where("user_id = ?", user_id).Find(rec)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no user records")
	}

	return rec, nil
}
