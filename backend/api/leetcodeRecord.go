package api

import (
	"mirror/db"
	"mirror/spider"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLeetcodeRecord(c *gin.Context) {
	val, found := db.GetLeetcodeRecordFromRedis(3)
	if found {
		c.String(http.StatusOK, val)
	} else {
		setting, has := db.GetLeetcodeSetting(3)
		if has == false {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Setting"})
			return
		}

		leetcodeSpider := spider.NewSpider("https://leetcode.com/api/user_submission_calendar/"+setting.UserName1, setting.Cookie)

		heatMap, err := leetcodeSpider.DoRequest()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Account"})
		}

		db.SetLeetcodeRecordToRedis(3, heatMap)

		c.String(http.StatusOK, heatMap)
	}
}
