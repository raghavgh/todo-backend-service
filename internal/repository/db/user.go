package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"todoapp/config"
	"todoapp/internal/models"
	"todoapp/internal/repository/errors"
)

type DBUserRepository struct {
	db *sql.DB
}

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

func NewDbUserRepository(cfg *config.TodoConfig) *DBUserRepository {
	dbUrl := cfg.DatabaseConfig.Url
	db, err := sql.Open("postgres", dbUrl)
	log.Println(dbUrl)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil
	}
	err = pingWithRetry(db, 5, time.Second*3)
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil
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
	return &DBUserRepository{
		db: db,
	}
}

func (d *DBUserRepository) GetUserByEmail(context context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users where email = $1`
	rows, err := d.db.Query(query, email)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.RepoErrors{Message: fmt.Sprintf("internal db error : %+v", err.Error())}
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return &user, err
}

func (d *DBUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, user_name, email, password) VALUES ($1, $2, $3, $4)`
	if d == nil {
		log.Printf("nil nil nil")
	}
	err := d.db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil
	}
	log.Println(user, query)
	if _, err := d.db.Exec(query, user.ID, user.UserName, user.Email, user.Password); err != nil {
		log.Println(err.Error())
		return errors.RepoErrors{Message: fmt.Sprintf("internal db error")}
	}
	return nil
}
