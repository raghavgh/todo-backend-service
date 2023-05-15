package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"todoapp/dto"
	"todoapp/internal/models"
	"todoapp/internal/repository"
	"todoapp/internal/repository/db"
	"todoapp/internal/repository/inmemory"
	"todoapp/utils"
)

type UserService struct {
	inMemoryRepo repository.UserRepositoryI
	dbRepo       repository.UserRepositoryI
	limit        int64
}

func NewUserService(inMemoryRepo *inmemory.InMemoryUserRepository, dbRepo *db.DBUserRepository) *UserService {
	return &UserService{
		inMemoryRepo: inMemoryRepo,
		dbRepo:       dbRepo,
	}
}

func (us *UserService) GetUserByEmail(context context.Context, email string) (user *models.User, err error) {
	if user, err = us.inMemoryRepo.GetUserByEmail(context, email); err != nil {
		user, err = us.dbRepo.GetUserByEmail(context, email)
		us.inMemoryRepo.CreateUser(context, user)
	}
	return user, err
}

func (us *UserService) CreateUser(ctx context.Context, signupRequest *dto.SignupRequest, id uuid.UUID) error {
	if signupRequest == nil {
		return ServiceErrors{message: fmt.Sprint("request is nil")}
	}
	user := &models.User{
		ID:       id,
		UserName: signupRequest.UserName,
		Email:    signupRequest.Email,
		Password: utils.GetHashPassword(signupRequest.Password),
	}
	err := us.dbRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	us.inMemoryRepo.CreateUser(ctx, user)
	return nil
}
