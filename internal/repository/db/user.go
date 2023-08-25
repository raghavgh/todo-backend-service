package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/raghavgh/todo-backend-service/internal/models"
	"github.com/raghavgh/todo-backend-service/internal/repository/errors"
)

type DBUserRepository struct {
	dbRepo *Repository
}

func NewDbUserRepository() *DBUserRepository {
	return &DBUserRepository{dbRepo: NewRepository()}
}

func (d *DBUserRepository) GetUserByEmail(context context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users where email = $1`
	rows, err := d.dbRepo.db.Query(query, email)
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
	err := d.dbRepo.db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil
	}
	log.Println(user, query)
	if _, err := d.dbRepo.db.Exec(query, user.ID, user.UserName, user.Email, user.Password); err != nil {
		log.Println(err.Error())
		return errors.RepoErrors{Message: fmt.Sprintf("internal db error")}
	}
	return nil
}
