package repository

import (
	"github.com/todo-backend-service/internal/repository/db"
	"github.com/todo-backend-service/internal/repository/inmemory"
)

func GetRepo(repoType string) UserRepositoryI {
	if repoType == "inMemory" {
		return inmemory.NewInMemoryUserRepository()
	}
	return db.NewDbUserRepository()
}
