package main

import (
	"errors"
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
		pongString := ""
		pongs, err := os.ReadFile(fmt.Sprintf("%s/pongs.log", logDir))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				pongString = "Could not find pongs.log file"
			} else {
				pongString = fmt.Sprintf("There was error reading pong count: %s", err)
			}
		}
		if pongString == "" {
			pongString = string(pongs)
		}
		htmlContent := fmt.Sprintf("<html><body><h1>Logs</h1><span style=\"white-space: pre-line\">%s</span><h1>Ping-pong count</h1><p>%s</p></body></html>", logsString, pongString)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
	router.Run()
}
