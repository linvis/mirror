package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type URLHandlers struct {
	URL      string
	Method   string
	Handlers gin.HandlersChain
}

// global url information
var URLHandle []URLHandlers

func RegisterURL(url string, method string, handlers ...gin.HandlerFunc) error {
	if method != "POST" && method != "GET" {
		return errors.New("method is not valid")
	}

	URLHandle = append(URLHandle, URLHandlers{url, method, handlers})

	return nil
}
