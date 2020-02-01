package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func BaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("mirror_token")
		if err != nil {
			log.Error("no cookie")
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Not Login"})
			c.Abort()
			return
		}
	}
}
