package controller

import (
	"fmt"
	"mirror/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type mockUser struct {
	Role   string `json:"role"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

// /user/login
func loginIn(c *gin.Context) {

	var info LoginInfo

	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
		return
	}

	fmt.Println(info)

	user := model.QueryUserByName("admin")
	fmt.Println(user)

	token := gin.H{"token": "admin-token"}

	// c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Account and password are incorrect."})
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": &token})
}

func userInfo(c *gin.Context) {
	token := c.PostForm("token")
	fmt.Println(token)

	user := mockUser{
		"admin",
		"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"Super Admin",
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": &user})
}

func InitUserRouter(engine *gin.Engine) {
	engine.POST("user/login", loginIn)
	engine.GET("user/info", userInfo)
}
