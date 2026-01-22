package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func Connect(user string, password string, host string, port string, database string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
	var db *sql.DB
	var err error
	tries := 1
	for tries <= 10 {
		fmt.Printf("Trying to connect to db, try %d ...\n", tries)
		dbLocal, errLocal := sql.Open("postgres", connStr)
		if errLocal == nil {
			if errLocal = dbLocal.Ping(); errLocal == nil {
				db = dbLocal
				break
			}
		}
		err = errLocal
		fmt.Printf("Error: %s\n", errLocal)
		fmt.Println("Failed connecting to db. Sleeping for 5 seconds before retry")
		time.Sleep(5 * time.Second)
		tries += 1
	}
	if db == nil {
		return nil, err
	}
	return db, nil
}

func Init(client *sql.DB) error {
	_, err := client.Query("CREATE TABLE IF NOT EXISTS todos (id TEXT PRIMARY KEY, content TEXT NOT NULL);")
	if err != nil {
		return err
	}
	return nil
}
