package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewMessage(content string) *Message {
	return &Message{
		ID:        uuid.New(),
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
