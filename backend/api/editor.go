package api

import (
	"mirror/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEditorCatalog(c *gin.Context) {

	id := GetUserID(c)

	catalog, err := db.GetCatalog(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Internal Error"})
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": catalog})
}
