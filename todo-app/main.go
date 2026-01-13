package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	fmt.Printf("Server started in port %s\n", port)

	router.GET("/", func(c *gin.Context) {
		htmlContent := "<html><body><h1>This is a great page :)</h1></body></html>"
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
	router.Run()
}
