package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/raghavgh/todo-backend-service/dto"
	"github.com/raghavgh/todo-backend-service/internal/models"
	"github.com/raghavgh/todo-backend-service/internal/repository"
	"github.com/raghavgh/todo-backend-service/internal/repository/db"
	"github.com/raghavgh/todo-backend-service/utils"
)

type UserService struct {
	dbRepo repository.UserRepositoryI
}

func NewUserService() *UserService {
	return &UserService{
		dbRepo: db.NewDbUserRepository(),
	}
}

func (us *UserService) GetUserByEmail(context context.Context, email string) (user *models.User, err error) {
	user, err = us.dbRepo.GetUserByEmail(context, email)
	if err != nil {
		return nil, err
	}
	return user, nil
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
	return nil
}
