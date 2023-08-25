package repository

import (
	"github.com/raghavgh/todo-backend-service/internal/repository/db"
	"github.com/raghavgh/todo-backend-service/internal/repository/inmemory"
)

func GetRepo(repoType string) UserRepositoryI {
	if repoType == "inMemory" {
		return inmemory.NewInMemoryUserRepository()
	}
	return db.NewDbUserRepository()
}
