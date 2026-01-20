package main

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	todos := []string{"test", "test2"}
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})
	router.POST("/todos", func(c *gin.Context) {
		todo := c.PostForm("todo")
		todo = strings.TrimSpace(todo)
		if len(todo) > 140 || len(todo) == 0 {
			c.JSON(400, gin.H{"error": "invalid-length"})
			return
		}
		todos = append(todos, todo)
		c.JSON(200, todo)
	})

	router.Run()
}
