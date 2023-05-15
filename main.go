package main

import (
	"log"
	"net/http"
	"todoapp/config"
	"todoapp/internal/server"

	valid "github.com/asaskevich/govalidator"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func init() {
	valid.SetFieldsRequiredByDefault(true)
	server.InitValidTags()
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %s\n", err.Error())
		panic(err)
	}
	// Create server
	server.Start(cfg)
}
