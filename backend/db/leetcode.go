package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

func SetLeetcodeRecordToRedis(user_id int, val interface{}) error {
	err := r.Set(strconv.Itoa(user_id)+"leetcode", val, time.Hour*24).Err()

	return err
}

func GetLeetcodeRecordFromRedis(user_id int) (string, bool) {
	found := false
	id := strconv.Itoa(user_id) + "leetcode"

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

type LeetcodeSetting struct {
	ID1       int    `xorm:"unsigned zerofill not null 'id'" json:"-"`
	UserID1   int    `xorm:"not null 'user_id'" json:"-"`
	UserName1 string `xorm:"not null 'user_name'" json:"leetcode_user_name"`
	Cookie    string `json:"cookie"`
}

func (s LeetcodeSetting) TableName() string {
	return "leetcode_setting"
}

func NewLeetcodeSetting(s *LeetcodeSetting) error {
	oldSetting := LeetcodeSetting{}

	has, err := x.Where("user_id = ?", s.UserID1).Get(&oldSetting)
	if err != nil {
		return err
	}

	if has {
		fmt.Println("update old leetcode settings", oldSetting)
		x.Update(s, &LeetcodeSetting{ID1: oldSetting.ID1})
	} else {
		_, err = x.Insert(s)
	}

	return err
}

func GetLeetcodeSetting(user_id int) (LeetcodeSetting, bool) {
	s := LeetcodeSetting{}

	has, _ := x.Where("user_id = ?", user_id).Get(&s)

	return s, has
}
