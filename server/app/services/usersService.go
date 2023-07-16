package services

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

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

func (s *UserService) UserAuthenticate(username, password string) (*models.UserRead, error) {
	user, err := models.UserModel().FindField("username", username)
	if err != nil {
		return &models.UserRead{}, err
	}

	if user == nil {
		return &models.UserRead{}, errors.New("User not found")
	}

	if !s.comparePasswords(user.Password, password) {
		return &models.UserRead{}, errors.New("Password is incorrect")
	}

	return user, nil
}

func (s *UserService) comparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func (s *UserService) CreateUser(username, email, password string) (*models.UserRead, error) {
	userModel := models.UserModel()
	user, _ := userModel.FindField("email", email)

	if user != nil {
		return &models.UserRead{}, errors.New("user already exist")
	}

	// Hash from password
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	data := models.UserWrite{
		Email:    email,
		Password: string(hashBytes),
		Username: username,
	}

	return userModel.Create(data)

}