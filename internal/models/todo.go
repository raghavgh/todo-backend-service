package models

import "time"

type Todo struct {
	Id          string    `json:"id" valid:"not_empty"`
	Message     string    `json:"message" valid:"not_empty"`
	Heading     string    `json:"heading" valid:"not_empty"`
	CreatedAt   time.Time `json:"createdAt" valid:"not_empty"`
	UpdatedAt   time.Time `json:"updatedAt" valid:"-"`
	CompletedAt time.Time `json:"completedAt" valid:"-"`
}
