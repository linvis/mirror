package api

import (
	"mirror/db"
	"mirror/spider"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLeetcodeRecord(c *gin.Context) {
	id := GetUserID(c)
	val, found := db.GetLeetcodeRecordFromRedis(id)
	if found {
		c.String(http.StatusOK, val)
	} else {
		setting, has := db.GetLeetcodeSetting(id)
		if has == false {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Setting"})
			return
		}

		leetcodeSpider := spider.NewSpider("https://leetcode.com/api/user_submission_calendar/"+setting.UserName1, setting.Cookie)

		heatMap, err := leetcodeSpider.DoRequest()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Account"})
		}

		db.SetLeetcodeRecordToRedis(id, heatMap)

		c.String(http.StatusOK, heatMap)
	}
}
