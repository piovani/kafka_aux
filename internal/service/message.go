package service

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/piovani/kafka_aux/internal/dto"
	"github.com/piovani/kafka_aux/internal/model"
	"github.com/piovani/kafka_aux/pkg/cache"
	"github.com/piovani/kafka_aux/pkg/queue"
)

type MessageService struct {
	cache *cache.Cache
	queue *queue.Queue
}

func NewMessageService() *MessageService {
	return &MessageService{
		cache: cache.NewCache(),
		queue: queue.NewQueue(),
	}
}

func (s *MessageService) All() ([]dto.MessageDto, error) {
	key, err := s.cache.GetAllKeys()
	if err != nil {
		return nil, err
	}

	all := []dto.MessageDto{}

	for _, key := range key {
		messageJson, err := s.cache.Get(key)
		if err != nil {
			fmt.Println(err)
		}

		var message model.Message
		if err = json.Unmarshal([]byte(messageJson), &message); err != nil {
			fmt.Println(err)
		}

		id, _ := uuid.Parse(key)
		var messageDto dto.MessageDto
		messageDto.ID = id
		messageDto.Content = message.Content

		all = append(all, messageDto)
	}

	return all, nil
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

	if err := s.queue.Publish(message.Content); err != nil {
		return messageDto, err
	}

	return messageDto, nil
}
