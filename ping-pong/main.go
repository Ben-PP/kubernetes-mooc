package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  	counter := 0
  	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
  	router.GET("/", func(c *gin.Context) {
  	    c.JSON(http.StatusOK, gin.H{"name": "ping-pong app"})
  	})
  	router.GET("/pingpong", func(c *gin.Context) {
  	    counter += 1
  	    c.HTML(http.StatusOK, "pingpong.tmpl", counter)
  	})
	router.GET("/pings", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"pings": counter})
	})
  	router.Run()
}
