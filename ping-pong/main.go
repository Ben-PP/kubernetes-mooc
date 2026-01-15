package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  counter := 0
  router := gin.Default()
  router.GET("/pingpong", func(c *gin.Context) {
	  htmlContent := fmt.Sprintf("<html><body><h1>pong %d</h1></body></html>", counter)
	  c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	  counter = counter + 1
  })
  router.Run()
}
