package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/kafka_aux/internal/dto"
	"github.com/piovani/kafka_aux/internal/service"
)

const (
	ErrInputCreateMessage      = "Unable to find message data"
	ErrValidInputCreateMessage = "Body invalid"
)

type MessageController struct {
	messageService service.MessageServiceContract
}

func NewMessageController() *MessageController {
	return &MessageController{
		messageService: service.NewMessageService(),
	}
}

func (c *MessageController) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (c *MessageController) Create(ctx *gin.Context) {
	messageDto, err := c.getMessageDto(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.GetErrorOutput(ErrInputCreateMessage))
		return
	}

	if err = c.validMessageDto(messageDto); err != nil {
		ctx.JSON(http.StatusForbidden, dto.GetErrorOutput(err.Error()))
		return
	}

	messageDto, err = c.messageService.Create(messageDto)
	if err != nil {
		ctx.JSON(http.StatusForbidden, dto.GetErrorOutput(ErrValidInputCreateMessage))
		return
	}

	ctx.JSON(http.StatusCreated, messageDto)
}

func (c *MessageController) getMessageDto(ctx *gin.Context) (dto.MessageDto, error) {
	var messageDto dto.MessageDto
	if err := ctx.BindJSON(&messageDto); err != nil {
		return messageDto, err
	}
	return messageDto, nil
}

func (c *MessageController) validMessageDto(msg dto.MessageDto) error {
	if msg.Content == "" {
		return fmt.Errorf("content is requred")
	}
	return nil
}
