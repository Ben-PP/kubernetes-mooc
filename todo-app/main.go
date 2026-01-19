package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"todo-app/directory"
	"todo-app/image"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	ImageURI template.URL
	Todos	 []string
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
	err := directory.MustExist(dataDir)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	fmt.Printf("Server started in port %s\n", port)
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		imageUri, err := image.GetImage(dataDir)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "internal-error.tmpl", err)
			return
		}
		data := PageData{
			ImageURI: template.URL(imageUri),
			Todos: []string{"test1", "test2"},
		}
		c.HTML(http.StatusOK, "index.tmpl", data) 
	})
	router.Run()
}
