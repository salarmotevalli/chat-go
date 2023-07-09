package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`
	AvatarImage string `json:"avatarImage" bson:"avatarImage"`
}

var userInstance User
var users *mongo.Collection

func UserModel() User {
	if users == nil {
		users = DB.Collection("users")
	}

	return userInstance
}

func (u *User) All() ([]*User, error) {
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

	cur.Close(Ctx)

	return result, nil
}
