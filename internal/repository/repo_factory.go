package repository

import (
	"todoapp/internal/repository/db"
	"todoapp/internal/repository/inmemory"
)

func GetRepo(repoType string) UserRepositoryI {
	if repoType == "inMemory" {
		return inmemory.NewInMemoryUserRepository()
	}
	return db.NewDbUserRepository()
}
