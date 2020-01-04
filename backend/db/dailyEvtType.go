package db

import "errors"

import "fmt"

type DailyEvtType struct {
	EvtType int    `xorm:"not null 'evt_type'" json:"evt_type"`
	EvtName string `xorm:"char(64) not null evt_name" json:"evt_name"`
	UserID  int    `xorm:"not null user_id" json:"user_id"`
}

func (evt DailyEvtType) TableName() string {
	return "daily_evt_type"
}

func CreateDailyEvtType(evt_type *DailyEvtType) error {
	_, err := x.Insert(evt_type)

	return err
}

func GetDailyEvtType(user_id int) ([]DailyEvtType, error) {

	allEvtTypes := []DailyEvtType{}

	err := x.Where("user_id = ?", user_id).Or("user_id = 0").Asc("evt_type").Find(&allEvtTypes)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no custom type")
	}

	fmt.Println(allEvtTypes)

	return allEvtTypes, nil
}

func DelDailyEvtType(evt_type int) error {
	_, err := x.Where("evt_type = ?", evt_type).Delete(&DailyEvtType{})
	if err != nil {
		return errors.New("delete evt fail")
	}

	return nil
}
