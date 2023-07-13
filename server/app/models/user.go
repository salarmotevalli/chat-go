package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRead struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	AvatarImage string             `json:"avatarImage" bson:"avatarImage"`
}

type UserWrite struct {
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`
	AvatarImage string `json:"avatarImage" bson:"avatarImage"`
}

type User struct{}

var userInstance User
var users *mongo.Collection

func UserModel() User {
	if users == nil {
		users = DB.Collection("users")
	}

	return userInstance
}

func (u User) All() ([]*UserRead, error) {
	query := bson.D{}
	usrs, err := Where(users, query, UserRead{})

	return usrs, err
}

func (u User) WhereEq(field string, target any) ([]*UserRead, error) {
	query := bson.M{field: bson.M{"$eq": target}}
	usrs, err := Where(users, query, UserRead{})

	return usrs, err
}

func (u User) WhereNe(field string, target any) ([]*UserRead, error) {
	query := bson.M{field: bson.M{"$ne": target}}
	usrs, err := Where(users, query, UserRead{})

	return usrs, err
}

func (u User) FindId(_id primitive.ObjectID) (any, error) {
	var result UserRead
	query := bson.D{bson.E{Key: "_id", Value: _id}}
	cur := users.FindOne(Ctx, query)

	err := cur.Decode(&result)
	if err != nil {
		return nil, err
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (_ User) Update(data map[string]interface{}, id interface{}) error {
	_, err := messages.UpdateByID(Ctx, id, data)

	return err
}
