package db

import "errors"

import "fmt"

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
