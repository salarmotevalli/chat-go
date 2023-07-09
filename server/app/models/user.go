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
	var result []*User

	// Run query
	cur, err := users.Find(Ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	// Collect data
	for cur.Next(Ctx) {
		var user User

		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		result = append(result, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	err = cur.Close(Ctx)
	if err != nil {
		return nil, err
	}

	var interfaces []interface{}
	for _, item := range result {
		interfaces = append(interfaces, interface{}(item))
	}
	return interfaces, nil
}

func (u UserQuery) WhereEq(field string, target any) ([]interface{}, error) {
	var result []*User
	query := bson.D{bson.E{Key: field, Value: target}}
	cur, err := users.Find(Ctx, query)
	if err != nil {
		return nil, err
	}

	// Collect data
	for cur.Next(Ctx) {
		var user User

		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		result = append(result, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	err = cur.Close(Ctx)
	if err != nil {
		return nil, err
	}

	var interfaces []interface{}
	for _, item := range result {
		interfaces = append(interfaces, interface{}(item))
	}
	return interfaces, nil
}

func (u UserQuery) Find(_id primitive.ObjectID) (any, error) {
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
