package db

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v7"
)

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

type DailyEvtItem struct {
	EvtItem     int    `xorm:"not null 'evt_item'" json:"evt_item"`
	EvtItemName string `xorm:"char(60) not null 'evt_item_name'" json:"evt_item_name"`
	EvtType     int    `xorm:"not null 'evt_type'" json:"evt_type"`
	UserID      int    `xorm:"not null 'user_id'" json:"-"`
}

func (item *DailyEvtItem) TableName() string {
	return "daily_evt_item"
}

func CreateDailyEvtItem(item *DailyEvtItem) error {
	_, err := x.Insert(item)

	return err
}

func GetDailyEvtItem(user_id int) ([]DailyEvtItem, error) {
	items := []DailyEvtItem{}

	err := x.Where("user_id = ?", user_id).Asc("evt_type").Find(&items)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no user custom items")
	}

	return items, nil
}

func DelDailyEvtItem(item int) error {
	_, err := x.Where("evt_item = ?", item).Delete(&DailyEvtItem{})
	if err != nil {
		return errors.New("delete evt fail")
	}

	return nil
}

func SetDailyEvtItemToRedis(user_id int, val string) error {

	err := r.Set(strconv.Itoa(user_id), val, 0).Err()

	return err
}

func GetDailyEvtItemFromRedis(user_id int) (string, bool) {
	found := false
	id := strconv.Itoa(user_id)

	val, err := r.Get(id).Result()

	if err == redis.Nil {
		// not found
	} else if err != nil {
		// error
	} else {
		found = true
	}

	return val, found
}
