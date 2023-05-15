package repository

import (
	"todoapp/config"
	"todoapp/internal/repository/db"
	"todoapp/internal/repository/inmemory"
)

func GetRepo(repoType string, cfg *config.TodoConfig) UserRepositoryI {
	if repoType == "inMemory" {
		return inmemory.NewInMemoryUserRepository(cfg)
	}
	return db.NewDbUserRepository(cfg)
}
