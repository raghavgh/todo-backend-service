package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"log"
	"net/http"
	"todoapp/config"
	"todoapp/dto"
	"todoapp/internal/models"
	"todoapp/internal/repository"
	"todoapp/internal/repository/db"
	"todoapp/internal/repository/inmemory"
	"todoapp/internal/services"
	"todoapp/utils"
)

type UserHandler struct {
	cfg     *config.TodoConfig
	service *services.UserService
}

func NewUserHandler(cfg *config.TodoConfig) *UserHandler {
	inMemoryRepo := repository.GetRepo("inMemory", cfg).(*inmemory.InMemoryUserRepository)
	dbRepo := repository.GetRepo("db", cfg).(*db.DBUserRepository)
	return &UserHandler{
		cfg:     cfg,
		service: services.NewUserService(inMemoryRepo, dbRepo),
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

// Signup handles all signup requests and create users
func (u *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest dto.SignupRequest
	// parsing request
	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	valid, err := govalidator.ValidateStruct(signupRequest)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}
	if !valid {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	id, _ := uuid.NewUUID()
	err = u.service.CreateUser(r.Context(), &signupRequest, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var token string
	// TODO Add a logic of refresh tokens in both backend and fronted for complete security
	token, err = utils.GetJWTToken(id, signupRequest.UserName, signupRequest.Email)
	json.NewEncoder(w).Encode(dto.SignupResponse{
		JwtToken: token,
		UserName: signupRequest.UserName,
		Email:    signupRequest.Email,
	})
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Printf("json error : %+v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid, err := govalidator.ValidateStruct(loginRequest)
	if err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	if !valid {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	var user *models.User
	user, err = u.service.GetUserByEmail(context.Background(), loginRequest.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal server error: %+v", err.Error()), http.StatusInternalServerError)
		return
	}
	if !utils.VerifyPassword(loginRequest.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var token string
	token, err = utils.GetJWTToken(user.ID, user.UserName, user.Email)
	json.NewEncoder(w).Encode(&dto.LoginResponse{
		UserName: user.UserName,
		Email:    user.Email,
		Todos:    nil,
		JwtToken: token,
	})
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var logoutRequest dto.LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&logoutRequest)
	if err != nil {
		log.Printf("json error : %+v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid, err := govalidator.ValidateStruct(logoutRequest)
	if err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	if !valid {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

}