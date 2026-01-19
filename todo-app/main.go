package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	ImageURI template.URL
}

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

func getLastUpdated(dir string) int64 {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	if len(entries) == 0 {
		return int64(0)
	}
	fmt.Println(entries)
	lastImage := entries[len(entries)-1].Name()
	lastUpdated := lastImage[:len(lastImage)-4]
	value, err := strconv.ParseInt(lastUpdated, 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func makeDataUri(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	enc := base64.StdEncoding.EncodeToString(b)
	uri := fmt.Sprintf("data:image/jpeg;base64,%s", enc)
	return uri, nil
}

func getImage(dataDir string) string {
	imageDir := fmt.Sprintf("%s/images", dataDir)
	dirExists, err := exists(imageDir)
	if err != nil {
		panic(err)
	}
	if !dirExists {
		err := os.MkdirAll(imageDir, 0777)
		if err != nil {
			panic(err)
		}
	}
	timestamp := time.Now().Unix()
	var lastUpdated int64 = getLastUpdated(imageDir)
	var fileName string = fmt.Sprintf("%s/%d.jpg", imageDir, lastUpdated)
	timeDif := timestamp - lastUpdated
	if timeDif > 600{

		fmt.Println(lastUpdated)
		resp, err := http.Get("https://picsum.photos/1200")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			panic(resp.StatusCode)
		}
		file, err := os.Create(fmt.Sprintf("%s/%d.jpg",imageDir, timestamp))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
		fileName = fmt.Sprintf("%s/%d.jpg", imageDir, timestamp)
		if lastUpdated != 0 {
			err := os.Remove(fmt.Sprintf("%s/%d.jpg", imageDir, lastUpdated))
			if err != nil {
				panic(err)
			}
		}
	}
	return fileName
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "./data"
	}
	fmt.Printf("DATA_DIR: %s", dataDir)
	dataDirExists, err := exists(dataDir)
	if err != nil {
		panic(err)
	}
	if !dataDirExists {
		err := os.MkdirAll(dataDir, 0777)
		if err != nil {
			panic(err)
		}
	}

	router := gin.Default()
	fmt.Printf("Server started in port %s\n", port)
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		imageFile := getImage(dataDir)
		imageUri, err := makeDataUri(imageFile)
		if err != nil {
			panic(err)
		}
		data := PageData{ImageURI: template.URL(imageUri)}
		c.HTML(http.StatusOK, "index.tmpl", data) 
	})
	router.Run()
}
