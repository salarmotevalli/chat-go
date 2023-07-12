package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	AvatarImage string             `json:"avatarImage" bson:"avatarImage"`
}

type UserQuery struct{}

var userInstance UserQuery
var users *mongo.Collection

func UserModel() UserQuery {
	if users == nil {
		users = DB.Collection("users")
	}

	return userInstance
}

func (u UserQuery) All() ([]interface{}, error) {
	query := bson.D{}

	return Where(users, query)
}

func (u UserQuery) WhereEq(field string, target any) ([]interface{}, error) {
	query := bson.M{field: bson.M{"$eq": target}}

	return Where(users, query)
}

func (u UserQuery) WhereNe(field string, target any) ([]interface{}, error) {
	query := bson.M{field: bson.M{"$ne": target}}

	return Where(users, query)
}

func (u UserQuery) FindId(_id primitive.ObjectID) (any, error) {
	var result User
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
