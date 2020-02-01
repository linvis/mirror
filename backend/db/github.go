package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

func SetGithubRecordToRedis(user_id int, val interface{}) error {
	err := r.Set(strconv.Itoa(user_id)+"-github", val, time.Hour*24).Err()

	return err
}

func GetGithubRecordFromRedis(user_id int) (string, bool) {
	found := false
	id := strconv.Itoa(user_id) + "-github"

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

type GithubSetting struct {
	ID2       int    `xorm:"unsigned zerofill not null 'id'" json:"-"`
	UserID2   int    `xorm:"not null 'user_id'" json:"-"`
	UserName2 string `xorm:"not null 'user_name'" json:"github_user_name"`
}

func (s GithubSetting) TableName() string {
	return "github_setting"
}

func NewGithubSetting(s *GithubSetting) error {
	oldSetting := GithubSetting{}

	has, err := x.Where("user_id = ?", s.UserID2).Get(&oldSetting)
	if err != nil {
		return err
	}

	if has {
		fmt.Println("update old github setting", oldSetting)
		x.Update(s, &GithubSetting{ID2: oldSetting.ID2})
	} else {
		_, err = x.Insert(s)
	}

	return err
}

func GetGithubSetting(user_id int) (GithubSetting, bool) {
	s := GithubSetting{}

	has, _ := x.Where("user_id = ?", user_id).Get(&s)

	return s, has
}
