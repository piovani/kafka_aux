package service

import (
	"encoding/json"
	"fmt"

	"github.com/piovani/kafka_aux/internal/dto"
	"github.com/piovani/kafka_aux/internal/model"
	"github.com/piovani/kafka_aux/pkg/cache"
)

type MessageService struct {
	cache *cache.Cache
}

func NewMessageService() *MessageService {
	return &MessageService{
		cache: cache.NewCache(),
	}
}

func (s *MessageService) Create(messageDto dto.MessageDto) (dto.MessageDto, error) {
	message := model.NewMessage(messageDto.Content)

	messageJson, err := json.Marshal(message)
	if err != nil {
		return messageDto, err
	}

	err = s.cache.Set(message.ID.String(), fmt.Sprintf("%s", messageJson), 0)
	if err != nil {
		return messageDto, err
	}

	messageDto.ID = message.ID
	return messageDto, nil
}
