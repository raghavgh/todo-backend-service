package repository

import (
	"context"
	_ "github.com/lib/pq"
	"todoapp/internal/models"
)

type UserRepositoryI interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
}
