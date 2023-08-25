package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/raghavgh/todo-backend-service/config"
)

type Repository struct {
	db *sql.DB
}

var repo = &Repository{}
var repoOnce sync.Once

func pingWithRetry(db *sql.DB, maxAttempts int, sleep time.Duration) error {
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = db.Ping()
		if err == nil {
			return nil
		}

		log.Printf("Attempt %d: Failed to ping the database: %v\n", attempt, err)

		// Sleep before the next attempt using exponential backoff
		time.Sleep(sleep * time.Duration(attempt))
	}

	return fmt.Errorf("failed to ping the database after %d attempts", maxAttempts)
}

func initDB() (*sql.DB, error) {
	dbUrl := config.Config.DatabaseConfig.Url
	var err error
	db := &sql.DB{}
	db, err = sql.Open("postgres", dbUrl)
	log.Println(dbUrl)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}
	err = pingWithRetry(db, 5, time.Second*3)
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}
	// Create a channel to receive signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to wait for signals and close the database connection
	go func() {
		s := <-sig
		log.Printf("Caught signal %v. Closing database connection...", s)
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()
	return db, err
}

func NewRepository() *Repository {
	repoOnce.Do(func() {
		db, err := initDB()
		if err != nil {
			log.Print(err)
		}
		repo.db = db
	})
	return repo
}
