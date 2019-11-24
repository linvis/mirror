package controller

import (
	"fmt"
	"mirror/model"
	"net/http"
	"regexp"

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
func loginIn(c *gin.Context) {

	var info LoginInfo

	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid Account"})
		return
	}

	fmt.Println(info)

	user := model.QueryUserByName(info.Username)
	if user.Password != info.Password {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
		return
	}

	token := gin.H{"token": "admin-token"}

	// c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": &token})
}

func userInfo(c *gin.Context) {
	token := c.Query("token")

	re := regexp.MustCompile(`(.*)-token`)
	name := re.FindStringSubmatch(token)
	fmt.Println(token, name)

	user := model.QueryUserByName(name[1])

	userInfo := UserInfo{
		user.Name,
		user.Avatar,
		user.Role,
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": &userInfo})
}

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "success"})
}

func init() {
	RegisterURL("user/login", "POST", loginIn)
	RegisterURL("user/info", "GET", userInfo)
	RegisterURL("user/logout", "POST", logout)
}
