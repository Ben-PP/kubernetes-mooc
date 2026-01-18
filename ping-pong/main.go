package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func logCount(logDir string, count int) {
	if err := os.WriteFile(fmt.Sprintf("%s/pongs.log", logDir), []byte(fmt.Sprintf("%d", count)), 0644); err != nil {
		panic(err)
	}
}

func main() {
  counter := 0
  logDir := os.Getenv("LOG_DIR")
  if logDir == "" {
	  logDir = "./logs"
  }
  fmt.Printf("LOG_DIR: %s", logDir)
  logDirExists, err := exists(logDir)
  if err != nil {
	  fmt.Printf("Error with logdir: %s", err)
	  return
  }
  if !logDirExists {
	  os.Mkdir(logDir, os.ModePerm)
  }
  
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{"name": "ping-pong app"})
  })
  router.GET("/pingpong", func(c *gin.Context) {
	  counter += 1
	  if err != nil {
		  panic(err)
	  }
	  htmlContent := fmt.Sprintf("<html><body><h1>pong %d</h1></body></html>", counter)
	  c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	  logCount(logDir, counter)
  })
  router.Run()
}
