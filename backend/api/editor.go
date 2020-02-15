package api

import (
	"fmt"
	"mirror/db"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetEditorCatalog(c *gin.Context) {

	id := GetUserID(c)

	catalog, err := db.GetCatalog(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Internal Error"})
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": catalog})
}

func NewEditorCatalog(c *gin.Context) {
	var catalog db.Catalog

	id := GetUserID(c)

	if err := c.BindJSON(&catalog); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid submit"})
		return
	}

	catalog.UserID = id

	log.Info("insert new catalog", catalog)

	err := db.NewCatalog(catalog)
	if err != nil {
		log.Warn(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid submit"})
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}
