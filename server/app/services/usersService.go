package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"chat/app/models"
)

type UserService struct{}

func LoadUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetOtherUsers(id string) ([]*models.UserRead, error) {
	// Convertr string id to object id
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// get users exept id
	userModel := models.UserModel()
	users, err := userModel.WhereNe("_id", _id)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (s *UserService) UpdateUser(data map[string]any, id string) error {
	// Convertr string id to object id
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Update model
	err = models.UserModel().Update(data, _id)

	return err
}


func (s *UserService) UpdateUserAvatar(avatar, id string) error {
	// Update model
	return s.UpdateUser(map[string]any{
		"avatarImage": avatar,
	}, id)
}
