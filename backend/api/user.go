package api

import (
	"fmt"
	"mirror/db"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Role   string `json:"role"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

// /user/login
func LoginIn(c *gin.Context) {

	var info LoginInfo

	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	fmt.Println(info)

	user, err := db.GetUserByName(info.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "internal errro"})
		return
	}
	if user.Password != info.Password {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
		return
	}

	token := gin.H{"token": strconv.Itoa(user.UserID)}

	// c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": &token})
}

func GetUserInfo(c *gin.Context) {
	token := c.Query("token")

	re := regexp.MustCompile(`(.*)-token`)
	name := re.FindStringSubmatch(token)
	fmt.Println(token, name)

	user, _ := db.GetUserByName(name[1])

	userInfo := UserInfo{
		user.Name,
		user.Avatar,
		user.Role,
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": &userInfo})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "success"})
}

func GetUserID(c *gin.Context) int {
	cookie, _ := c.Cookie("mirror_token")
	id, _ := strconv.Atoi(cookie)

	return id
}
