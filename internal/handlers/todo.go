package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/raghavgh/todo-backend-service/dto"
	"github.com/raghavgh/todo-backend-service/internal/services"
	"github.com/raghavgh/todo-backend-service/utils"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		service: services.NewTodoService(),
	}
}

func (t *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (t *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	createTodoRequest := &dto.CreateTodoRequest{}
	err := json.NewDecoder(r.Body).Decode(createTodoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := utils.GetUserIDFromJWT(createTodoRequest.JwtToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = t.service.CreateTodo(context.Background(), createTodoRequest, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
