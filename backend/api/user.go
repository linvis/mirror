package api

import (
	"fmt"
	"mirror/db"
	"mirror/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
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
		log.Warnf("login fail, input pw:%s, need pw: %s", user.Password, info.Password)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
		return
	}

	// token := gin.H{"token": strconv.Itoa(user.UserID)}
	token := middleware.SetToken(info.Username, info.Password, user.UserID)

	c.SetCookie("mirror", token, 24*3600, "/", "localhost", false, true)

	// c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "success"})
}

func GetUserInfo(c *gin.Context) {

	id := GetUserID(c)

	user, _ := db.GetUserByID(id)

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

	val, _ := c.Get("userID")

	return val.(int)
}
