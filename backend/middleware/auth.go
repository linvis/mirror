package middleware

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"io"
	"mirror/db"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type authSalt struct {
	Salt string `json:"salt"`
}

var salt authSalt

func BaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		handleFail := func() {
			log.Error("no cookie")
			c.JSON(http.StatusOK, gin.H{"code": 60204, "message": "Not Login"})
			c.Abort()
		}

		token, err := c.Cookie("mirror")
		if err != nil {
			handleFail()
			return
		}

		id, err := GetUserIdFromToken(token)
		if err != nil {
			handleFail()
			return
		}
		log.Debug("middle auth", token, id)
		c.Set("userID", id)
	}
}

func SetToken(user string, password string, id int) string {
	h := md5.New()
	io.WriteString(h, user)
	io.WriteString(h, password)
	io.WriteString(h, salt.Salt)

	token := string(h.Sum(nil))

	err := db.SetToRedis(token, strconv.Itoa(id), time.Hour*24)
	if err != nil {
		log.Warn("set token to redis fail")
	}

	return token
}

func GetUserIdFromToken(token string) (int, error) {
	val, err := db.GetFromRedis(token)
	if err != nil {
		return -1, errors.New("invalid token")
	}

	id, _ := strconv.Atoi(val)

	return id, nil
}

func init() {

	file, err := os.Open("./middleware/salt.config")
	defer file.Close()

	if err != nil {
		log.Warn("open auth salt error", err)
		return
	}

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&salt)
}
