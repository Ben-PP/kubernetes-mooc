package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"todo-backend/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func main() {
	log.SetOutput(os.Stdout)
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbDatabase := os.Getenv("POSTGRES_DATABASE")
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbDatabase == "" {
		log.Fatal("Postgres env variables must be defined")
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	dbClient, err := db.Connect(dbUser, dbPassword, dbHost, dbPort, dbDatabase)
	if err != nil {
		log.Fatal("Failed to connect to db: %s", err)
	}
	defer dbClient.Close()
	if err := db.Init(dbClient); err != nil {
		log.Fatal("Failed to init db: %s", err)
	}
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", func(c *gin.Context) {
		rows, err := dbClient.Query("SELECT content FROM todos;")
		if err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to query todos: %s", err)})
			return
		}
		defer rows.Close()
		var todos []string
		for rows.Next() {
			var content string
			if err := rows.Scan(&content); err != nil {
				c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to scan rows: %s", err)})
				return
			}
			todos = append(todos, content)
		}
		c.JSON(http.StatusOK, todos)
	})
	router.POST("/todos", func(c *gin.Context) {
		todo := c.PostForm("todo")
		todo = strings.TrimSpace(todo)
		if len(todo) > 140 || len(todo) == 0 {
			log.Print(fmt.Sprintf("Blocked too long todo: %s", todo))
			c.JSON(400, gin.H{"error": "invalid-length"})
			return
		}
		todoID := uuid.New().String()
		_, err := dbClient.Query("INSERT INTO todos (id, content) VALUES ($1,$2);",todoID, todo)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		log.Print(fmt.Sprintf("Created todo: %s", todo))
		c.JSON(200, todo)
	})

	router.Run()
}
