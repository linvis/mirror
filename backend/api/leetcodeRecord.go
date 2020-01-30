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
		leetcodeCookies, err := spider.LoadCookies("spider/leetcode.cookie")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Cookies"})
		}

		leetcodeSpider := spider.NewSpider("https://leetcode.com/api/user_submission_calendar/slinz/", leetcodeCookies)

		heatMap, err := leetcodeSpider.DoRequest()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Cookies"})
		}

		db.SetLeetcodeRecordToRedis(3, heatMap)

		c.String(http.StatusOK, heatMap)
	}
}
