package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"todoapp/config"
	"todoapp/internal/models"
	"todoapp/internal/repository/errors"
)

var mu sync.RWMutex

type InMemoryUserRepository struct {
	usersByID    map[uuid.UUID]*models.User
	usersByEmail map[string]*models.User
}

func NewInMemoryUserRepository(cfg *config.TodoConfig) *InMemoryUserRepository {
	return &InMemoryUserRepository{
		usersByEmail: make(map[string]*models.User),
		usersByID:    make(map[uuid.UUID]*models.User),
	}
}

func (i *InMemoryUserRepository) GetUserByEmail(context context.Context, email string) (*models.User, error) {
	mu.RLock()
	defer mu.RUnlock()
	if user, ok := i.usersByEmail[email]; ok {
		return user, nil
	}
	return nil, errors.RepoErrors{Message: fmt.Sprintf("user with email %s not ound", email)}
}

func (i *InMemoryUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	mu.Lock()
	defer mu.Unlock()

	if user == nil {
		return errors.RepoErrors{Message: fmt.Sprint("user is nil")}
	}
	if _, ok := i.usersByEmail[user.Email]; ok {
		return errors.RepoErrors{Message: fmt.Sprintf("user with email %s is already exist", user.Email)}
	}
	if len(i.usersByEmail) > 100 {
		for key, _ := range i.usersByEmail {
			delete(i.usersByEmail, key)
			break
		}
	}
	i.usersByEmail[user.Email] = user
	i.usersByID[user.ID] = user
	fmt.Println(i.usersByID, i.usersByEmail)
	return nil
}
