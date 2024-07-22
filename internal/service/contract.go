package service

import "github.com/piovani/kafka_aux/internal/dto"

type MessageServiceContract interface {
	Create(messageDto dto.MessageDto) (dto.MessageDto, error)
}
