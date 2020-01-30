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
