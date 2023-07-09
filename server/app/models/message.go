package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Message string             `json:"message" bson:"message"`
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

func (m MessageQuery) All() ([]interface{}, error) {
	var result []*Message

	// Run query
	cur, err := messages.Find(Ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	// Collect data
	for cur.Next(Ctx) {
		var message Message

		err := cur.Decode(&message)
		if err != nil {
			return nil, err
		}

		result = append(result, &message)
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

func (m MessageQuery) WhereEq(field string, target any) ([]interface{}, error) {
	var result []*Message
	query := bson.D{bson.E{Key: field, Value: target}}
	cur, err := users.Find(Ctx, query)

	if err != nil {
		return nil, err
	}

	// Collect data
	for cur.Next(Ctx) {
		var message Message

		err := cur.Decode(&message)
		if err != nil {
			return nil, err
		}

		result = append(result, &message)
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

func (m MessageQuery) Find(_id primitive.ObjectID) (any, error) {
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
