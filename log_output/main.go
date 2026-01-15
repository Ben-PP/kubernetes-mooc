package main

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	randomString := id.String()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		loc, _ := time.LoadLocation("Europe/Helsinki")
		currentTime := time.Now().In(loc)
		statusLine := fmt.Sprintf("%s: %s", currentTime, randomString)
		htmlContent := fmt.Sprintf("<html><body><h1>Status</h1><p>%s</p></body></html>", statusLine)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
	router.Run()
}
