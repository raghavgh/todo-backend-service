package dto

import "todoapp/internal/models"

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
