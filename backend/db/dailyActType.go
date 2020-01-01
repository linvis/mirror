package db

import "errors"

import "fmt"

type DailyActType struct {
	ActType int    `xorm:"not null 'act_type'" json:"act_type"`
	ActName string `xorm:"char(64) not null act_name" json:"act_name"`
	UserID  int    `xorm:"user_id" json:"user_id"`
}

func (evt DailyActType) TableName() string {
	return "daily_act_type"
}

func GetUserActType(id int) (*[]DailyActType, error) {

	baseTypes := []DailyActType{}
	customTypes := []DailyActType{}

	err := x.Table("daily_act_type").Where("user_id = ?", 0).Find(&baseTypes)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no base type")
	}

	err = x.Table("daily_act_type").Where("user_id = ?", id).Find(&customTypes)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no custom type")
	}

	baseTypes = append(baseTypes, customTypes...)
	fmt.Println(baseTypes)

	return &baseTypes, nil
}
