package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"chat/app/models"
)

type MessageService struct{}

func LoadMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) CreateMessage(from, to, message string) error {
	messageModel := models.MessageModel()

	senderObjectId, err := primitive.ObjectIDFromHex(from)
	if err != nil {
		return err
	}

	var data models.MessageWrite = models.MessageWrite{
		Message: message,
		Users:   []string{from, to},
		Sender:  senderObjectId,
	}

	return messageModel.Create(data)
}

func (s *MessageService) GetMessages(from, to string) ([]*models.MessageRead, error)  {
	messageModel := models.MessageModel()

	messages, err := messageModel.WhereEq("users", []string{from, to,})
	if err != nil {
		return []*models.MessageRead{}, err
	}

	return messages, err
}