package main

import (
	"fmt"
	"log-server/pingpong"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type IndexPageData struct {
	Logs	string
	Pings	string
}

func main() {
	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "./logs"
	}
	fmt.Printf("LOG_DIR: %s\n", logDir)
	pingpongURL := os.Getenv("PINGPONG_URL")
	if pingpongURL == "" {
		panic("PINGPONG_URL can not be empty")
	}
	fmt.Printf("PINGPONG_URL: %s\n", pingpongURL)

	pongClient := pingpong.New(pingpongURL)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		logsString := ""
		logs, err := os.ReadFile(fmt.Sprintf("%s/timestamps.log", logDir))
		if err != nil {
			logsString = fmt.Sprintf("There was error getting logs: %s", err)
		}
		logsString = string(logs)
		pongString := ""
		pongs, err := pongClient.Pings()
		if err != nil {
			pongString = fmt.Sprintf("There was error getting pong count: %s", err)
		}
		if pongString == "" {
			pongString = fmt.Sprintf("%d", pongs)
		}
		data := IndexPageData{
			Logs: logsString,
			Pings: pongString,
		}
		c.HTML(http.StatusOK, "index.tmpl", data)
	})
	router.Run()
}
