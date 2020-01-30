package api

import (
	"fmt"
	"mirror/db"
	"mirror/spider"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLeetcodeRecord(c *gin.Context) {
	var rec db.LeetcodeRecord

	if err := c.BindJSON(&rec); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	fmt.Println(rec)

	rec.UserID = 3
	err := db.NewLeetcodeRecord(&rec)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

func GetLeetcodeRecord(c *gin.Context) {

	leetcodeCookies, err := spider.LoadCookies("spider/leetcode.cookie")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Cookies"})
	}

	leetcodeSpider := spider.NewSpider("https://leetcode.com/api/user_submission_calendar/slinz/", leetcodeCookies)

	heatMap, err := leetcodeSpider.DoSpider()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Leetcode Cookies"})
	}

	// c.JSON(http.StatusOK, gin.H{"code": 20000, "data": string(heatMap)})
	c.String(http.StatusOK, heatMap)
	// c.Data(http.StatusOK, "json", heatMap)
}
