package api

import (
	"encoding/json"
	"mirror/db"
	"mirror/spider"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetGithubRecord(c *gin.Context) {

	val, found := db.GetGithubRecordFromRedis(3)
	if found {
		c.String(http.StatusOK, val)
	} else {
		setting, has := db.GetGithubSetting(3)
		if has == false {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Github setting"})
			return
		}
		spider := spider.NewSpider("https://github.com/users/"+setting.UserName2+"/contributions", "")
		res, err := spider.DoRequest()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Github Account"})
		}

		r := regexp.MustCompile("data-count=\"(\\d)\" data-date=\"(\\d+-\\d+-\\d+)\"")
		match := r.FindAllStringSubmatch(res, -1)

		if len(match) <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Can't load github record"})
			return
		}

		records := make(map[string]int)

		for i := range match {
			// 0 2020-01-31
			count, _ := strconv.Atoi(match[i][1])

			if count <= 0 {
				continue
			}

			t, _ := time.Parse("2006-01-02", match[i][2])
			timestamp := strconv.FormatInt(t.Unix(), 10)

			records[timestamp] = count
		}

		b, _ := json.Marshal(records)

		b, _ = json.Marshal(string(b))

		db.SetGithubRecordToRedis(3, string(b))

		c.String(http.StatusOK, string(b))
	}
}
