package main

import (
	"todoapp/internal/server"

	valid "github.com/asaskevich/govalidator"
)

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

func main() {
	// Create server
	server.Start()
}
