package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Message   string             `json:"message" bson:"message"`
	Users     []string           `json:"users" bson:"users"`
	Sender    primitive.ObjectID `json:"sender" bson:"sender"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}

type MessageQuery struct{}

var messageInstance MessageQuery
var messages *mongo.Collection

func MessageModel() MessageQuery {
	if messages == nil {
		messages = DB.Collection("messages")
	}

	return messageInstance
}

func (m MessageQuery) All() ([]*Message, error) {
	msgs, err := Where(messages, bson.D{}, Message{})
	return msgs, err
}

func (m MessageQuery) WhereEq(field string, target any) ([]*Message, error) {
	query := bson.M{field: bson.M{"$all": target}}
	msgs, err := Where(messages, query, Message{})

	return msgs, err

}

func (m MessageQuery) FindId(_id primitive.ObjectID) (any, error) {
	var result Message
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
