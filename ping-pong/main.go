package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func initDB(db *sql.DB) {
	_, err := db.Query("CREATE TABLE IF NOT EXISTS pings (id INTEGER PRIMARY KEY, count INTEGER);")
	if err != nil {
		panic(err)
	}
	var count int
	err = db.QueryRow("SELECT count FROM pings WHERE id=1;").Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == sql.ErrNoRows {
		_, err := db.Query("INSERT INTO pings (id, count) VALUES (1,0);")
		if err != nil {
			panic(err)
		}
	}
}

func getCount(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT count FROM pings;").Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if err == sql.ErrNoRows {
		count = 0
	}
	return count, nil
}

func main() {
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	if pgHost == "" || pgPassword == "" {
		panic("POSTGRES_HOST and POSTGRES_PASSWORD must be provided")
	}
	if pgPort == "" {
		pgPort = "5432"
	}
	connStr := fmt.Sprintf("postgres://postgres:%s@%s:%s/postgres?sslmode=disable", pgPassword, pgHost, pgPort)
	var db *sql.DB
	tries := 1
	for tries <= 10 {
		fmt.Printf("Trying to connect to db, try %d ...\n", tries)
		dbLocal, err := sql.Open("postgres", connStr)
		if err == nil {
			db = dbLocal
			defer db.Close()
			break
		}
		fmt.Printf("Error: %s\n", err)
		fmt.Println("Failed connecting to db. Sleeping for 5 seconds before retry")
		time.Sleep(5 * time.Second)
		tries += 1
	}
	initDB(db)

  	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
  	router.GET("/", func(c *gin.Context) {
  	    c.JSON(http.StatusOK, gin.H{"name": "ping-pong app"})
  	})
  	router.GET("/pingpong", func(c *gin.Context) {
		var count int
		if err := db.QueryRow("SELECT count FROM pings WHERE id = 1;").Scan(&count); err != nil {
  	    	c.HTML(http.StatusOK, "pingpong.tmpl", err)
			return
		}
		_, err := db.Query("UPDATE pings SET count = count + 1 WHERE id = 1;")
		if err != nil {
  	    	c.HTML(http.StatusOK, "pingpong.tmpl", err)
		}
			
  	    c.HTML(http.StatusOK, "pingpong.tmpl", count+1)
  	})
	router.GET("/pings", func(c *gin.Context) {
		count, err := getCount(db)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pings": count})
	})
	router.GET("/healthz", func(c *gin.Context) {
		_, err := getCount(db)
		if err != nil {
			c.JSON(500, gin.H{"alive": 0})
			return
		}
		c.JSON(http.StatusOK, gin.H{"alive": 1})
	})
  	router.Run()
}
