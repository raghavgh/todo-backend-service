package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/raghavgh/gofast"
	"github.com/raghavgh/todo-backend-service/config"
	"github.com/raghavgh/todo-backend-service/internal/models"
	"github.com/raghavgh/todo-backend-service/internal/repository/errors"
)

var mu sync.RWMutex

type InMemoryUserRepository struct {
	emailCache gofast.Cache
	idCache    gofast.Cache
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		emailCache: gofast.NewCache(config.Config.CacheConfig.Limit, 1),
		idCache:    gofast.NewCache(config.Config.CacheConfig.Limit, 1),
	}
}

func (i *InMemoryUserRepository) GetUserByEmail(context context.Context, email string) (*models.User, error) {
	val, ok := i.emailCache.Get(email)
	if ok {
		return val.(*models.User), nil
	}
	return nil, errors.RepoErrors{Message: fmt.Sprintf("user with email %s not found", email)}
}

func (i *InMemoryUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	mu.Lock()
	defer mu.Unlock()

	if user == nil {
		return errors.RepoErrors{Message: fmt.Sprint("user is nil")}
	}
	if ok := i.emailCache.Contains(user.Email); ok {
		return errors.RepoErrors{Message: fmt.Sprintf("user with email %s is already exist", user.Email)}
	}
	i.emailCache.Put(user.Email, user)
	if ok := i.idCache.Contains(user.ID.String()); ok {
		return errors.RepoErrors{Message: fmt.Sprintf("user with id %s is already exist", user.ID)}
	}
	i.idCache.Put(user.ID.String(), user)
	return nil
}
