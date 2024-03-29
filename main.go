package main

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/raghavgh/todo-backend-service/internal/server"
)

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

func main() {
	// Create server
	server.Start()
}
