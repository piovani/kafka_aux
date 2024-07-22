package dto

import "github.com/google/uuid"

type MessageDto struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}
