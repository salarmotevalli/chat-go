package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageOutput struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Message   string             `json:"message" bson:"message"`
	Users     []string           `json:"users" bson:"users"`
	Sender    primitive.ObjectID `json:"sender" bson:"sender"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}

type MessageInput struct {
	Message string             `json:"message" bson:"message"`
	Users   []string           `json:"users" bson:"users"`
	Sender  primitive.ObjectID `json:"sender" bson:"sender"`
}

type Message struct{}

var messageInstance Message
var messages *mongo.Collection

func MessageModel() Message {
	if messages == nil {
		messages = DB.Collection("messages")
	}

	return messageInstance
}

func (m Message) All() ([]*MessageOutput, error) {
	msgs, err := Where(messages, bson.D{}, MessageOutput{})
	return msgs, err
}

func (m Message) WhereEq(field string, target any) ([]*MessageOutput, error) {
	query := bson.M{field: bson.M{"$all": target}}
	msgs, err := Where(messages, query, MessageOutput{})

	return msgs, err

}

func (m Message) FindId(_id primitive.ObjectID) (any, error) {
	var result MessageOutput
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
