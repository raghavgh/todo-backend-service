package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
