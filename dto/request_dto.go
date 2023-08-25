package dto

import (
	"github.com/raghavgh/todo-backend-service/internal/epoch"
	"github.com/raghavgh/todo-backend-service/internal/models"
)

type LoginRequest struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"not_empty, password"`
}

type SignupRequest struct {
	UserName string `json:"userName" valid:"not_empty"`
	Email    string `json:"email" valid:"not_empty, email"`
	Password string `json:"password" valid:"not_empty, password"`
}

type SignupResponse struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	JwtToken string `json:"jwtToken"`
}

type LoginResponse struct {
	UserName string        `json:"userName" valid:"not_empty"`
	Email    string        `json:"email" valid:"not_empty, email"`
	JwtToken string        `json:"jwtToken" valid:"not_empty"`
	Todos    []models.Todo `json:"todos"`
}

type LogoutRequest struct {
	Email string `json:"email" valid:"not_empty"`
}

type CreateTodoRequest struct {
	JwtToken    string                `json:"jwtToken"`
	ID          string                `json:"id"`
	Heading     string                `json:"heading"`
	Message     string                `json:"message"`
	CreatedAt   epoch.EpochMillisTime `json:"createdAt"`
	UpdatedAt   epoch.EpochMillisTime `json:"updatedAt"`
	CompletedAt epoch.EpochMillisTime `json:"completedAt"`
}

type GetTodosRequest struct {
	JwtToken string `json:"jwtToken"`
}
