package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func init() {
	RegisterURL("/", "GET", home)
}
