package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("mirror_token")
		if err != nil {
			fmt.Println("no cookie")
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Not Login"})
			c.Abort()
			return
		}
	}
}
