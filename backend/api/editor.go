package api

import (
	"fmt"
	"mirror/db"
	"net/http"
	"strconv"

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

func UpdateEditorCatalog(c *gin.Context) {
	var catalog []db.Catalog

	id := GetUserID(c)

	if err := c.BindJSON(&catalog); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid submit"})
		return
	}

	catalog[0].UserID = id

	log.Info("insert new catalog", catalog)

	err := db.UpdateCatalog(catalog[0])
	if err != nil {
		log.Warn(err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid submit"})
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": ""})
}

func NewDocument(c *gin.Context) {
	var doc db.Document

	id := GetUserID(c)
	doc.UserID = id

	if err := c.BindJSON(&doc); err != nil {
		log.Warn("new doc invalid request", err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid request"})
		return
	}

	log.Debug(doc)
	err := db.NewDocument(&doc)
	if err != nil {
		log.Warn("new doc save fail", err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": gin.H{"id": doc.Id}})
}

func DeleteDocument(c *gin.Context) {

	docID := c.Param("docid")

	id, err := strconv.Atoi(docID)
	if err != nil {
		log.Error("invalid doc id ", err)
		c.JSON(http.StatusOK, gin.H{"code": 20000})
		return
	}

	err = db.DeleteDocument(id)
	if err != nil {
		log.Error("delete doc fail ", err)
		c.JSON(http.StatusOK, gin.H{"code": 20000})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000})
	return
}

func GetAllDocument(c *gin.Context) {

	id := GetUserID(c)

	docs, err := db.GetAllDocument(id)
	if err != nil {
		log.Warn("get doc fail", err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": docs})
}

func GetDocumentByID(c *gin.Context) {

	user_id := GetUserID(c)
	doc_id := c.Param("docid")

	docs, err := db.GetDocumentByID(user_id, doc_id)
	if err != nil {
		log.Warn("get doc fail", err)
		c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": docs})
}
