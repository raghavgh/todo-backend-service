package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/raghavgh/todo-backend-service/dto"
	"github.com/raghavgh/todo-backend-service/internal/models"
	"github.com/raghavgh/todo-backend-service/internal/repository"
	"github.com/raghavgh/todo-backend-service/internal/repository/db"
)

type TodoService struct {
	dbRepo repository.TodoRepositoryI
}

func NewTodoService() *TodoService {
	return &TodoService{dbRepo: db.NewDbTodoRepository()}
}

func (ts *TodoService) CreateTodo(ctx context.Context, createTodoRequest *dto.CreateTodoRequest, userID string) error {
	if createTodoRequest == nil {
		return errors.New("request body cannot be empty")
	}

	todo := &models.Todo{
		ID:          createTodoRequest.ID,
		Message:     createTodoRequest.Message,
		Heading:     createTodoRequest.Heading,
		CreatedAt:   time.Time(createTodoRequest.CreatedAt),
		UpdatedAt:   time.Time(createTodoRequest.UpdatedAt),
		CompletedAt: time.Time(createTodoRequest.CompletedAt),
		UserID:      userID,
	}
	fmt.Println(todo)
	return nil
}
