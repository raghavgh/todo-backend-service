package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/raghavgh/todo-backend-service/config"
	"github.com/raghavgh/todo-backend-service/constants"
	"github.com/raghavgh/todo-backend-service/internal/handlers"
	"github.com/raghavgh/todo-backend-service/utils"
)

type Server http.Server

func NewServer(router *mux.Router) *Server {
	return &Server{
		Addr:         fmt.Sprintf(":%s", config.Config.ServerConfig.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(config.Config.ServerConfig.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.Config.ServerConfig.WriteTimeout) * time.Millisecond,
	}
}

func Start() {
	router := BuildRouter()
	s := NewServer(router)

	log.Printf("Starting new server at %s", s.Addr)
	go func() {
		if err := http.ListenAndServe(s.Addr, s.Handler); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	log.Println("Shutting down server...")
	log.Println("Server gracefully shut down")
}

// Enable CORS middleware
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func BuildRouter() *mux.Router {
	userHandler := handlers.NewUserHandler()
	todoHandler := handlers.NewTodoHandler()
	router := mux.NewRouter()

	router.Use(enableCors) // Add CORS middleware

	protectedRouter := router.PathPrefix("/api").Subrouter()
	protectedRouter.Use(utils.AuthMiddleware)
	router.HandleFunc(constants.SIGNUP, userHandler.Signup).Methods("POST")
	router.HandleFunc("/check", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}).Methods("GET")
	router.HandleFunc(constants.LOGIN, userHandler.Login).Methods("POST")

	protectedRouter.HandleFunc(constants.USER, userHandler.GetUser).Methods("POST")
	protectedRouter.HandleFunc(constants.CREATE_TODO, todoHandler.Create).Methods("POST")

	return router
}
