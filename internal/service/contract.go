package service

import (
	"github.com/piovani/kafka_aux/internal/dto"
)

type MessageServiceContract interface {
	All() ([]dto.MessageDto, error)
	Create(messageDto dto.MessageDto) (dto.MessageDto, error)
}
