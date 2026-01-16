package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "./logs"
	}
	fmt.Printf("LOG_DIR: %s\n", logDir)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		logsString := ""
		logs, err := os.ReadFile(fmt.Sprintf("%s/timestamps.log", logDir))
		if err != nil {
			logsString = fmt.Sprintf("There was error getting logs: %s", err)
		}
		logsString = string(logs)
		htmlContent := fmt.Sprintf("<html><body><h1>Logs</h1><span style=\"white-space: pre-line\">%s</span></body></html>", logsString)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
	router.Run()
}
