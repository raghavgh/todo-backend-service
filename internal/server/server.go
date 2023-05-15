package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	valid "github.com/asaskevich/govalidator"

	"todoapp/config"
	"todoapp/constants"
	"todoapp/internal/handlers"
	"todoapp/utils"
)

func InitValidTags() {
	valid.TagMap["not_empty"] = func(str string) bool {
		return len(str) != 0
	}
	valid.TagMap["password"] = func(password string) bool {
		if valid.IsNull(password) {
			return false
		}
		if len(password) < 8 {
			return false
		}
		if !valid.StringMatches(password, `[A-Z]`) {
			return false
		}
		if !valid.StringMatches(password, `[a-z]`) {
			return false
		}
		if !valid.StringMatches(password, `[0-9]`) {
			return false
		}
		return true
	}
}

type Server http.Server

func NewServer(router *mux.Router, cfg *config.TodoConfig) *Server {
	return &Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeout) * time.Millisecond,
	}
}

func Start(cfg *config.TodoConfig) {

	router := BuildRouter(cfg)
	s := NewServer(router, cfg)

	log.Printf("Starting server at %s", s.Addr)
	go func() {
		if err := http.ListenAndServe(s.Addr, s.Handler); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Listen for SIGINT and SIGTERM signals and shut down the server gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	log.Println("Shutting down server...")

	log.Println("Server gracefully shut down")

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func BuildRouter(cfg *config.TodoConfig) *mux.Router {
	userHandler := handlers.NewUserHandler(cfg)
	router := mux.NewRouter()
	protectedRouter := router.PathPrefix("/api").Subrouter()
	protectedRouter.Use(utils.AuthMiddleware)
	router.HandleFunc(constants.SIGNUP, userHandler.Signup).Methods("POST")
	router.HandleFunc("/check", func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)
		writer.WriteHeader(http.StatusOK)
	}).Methods("GET")
	router.HandleFunc(constants.LOGIN, userHandler.Login).Methods("POST")
	protectedRouter.HandleFunc(constants.USER, userHandler.GetUser).Methods("POST")
	return router
}
