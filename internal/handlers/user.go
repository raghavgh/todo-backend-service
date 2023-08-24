package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todoapp/dto"
	"todoapp/internal/models"
	"todoapp/internal/services"
	"todoapp/internal/validators"
	"todoapp/utils"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		service: services.NewUserService(),
	}
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

// Signup handles all signup requests and create users
func (u *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest *dto.SignupRequest
	// parsing request
	err := json.NewDecoder(r.Body).Decode(signupRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	valid, err := validators.IsValidSignupRequest(signupRequest)
	if !valid {
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}
	id, _ := uuid.NewUUID()
	err = u.service.CreateUser(r.Context(), signupRequest, id)
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
	var loginRequest *dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(loginRequest)
	if err != nil {
		log.Printf("json error : %+v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid, err := validators.IsValidLoginRequest(loginRequest)
	if !valid {
		http.Error(w, "Validation error", http.StatusBadRequest)
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
