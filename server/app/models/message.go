package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRead struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Message   string             `json:"message" bson:"message"`
	Users     []string           `json:"users" bson:"users"`
	Sender    primitive.ObjectID `json:"sender" bson:"sender"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type MessageWrite struct {
	Message   string             `json:"message" bson:"message"`
	Users     []string           `json:"users" bson:"users"`
	Sender    primitive.ObjectID `json:"sender" bson:"sender"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
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

func (m Message) All() ([]*MessageRead, error) {
	msgs, err := Where(messages, bson.D{}, MessageRead{})
	return msgs, err
}

func (m Message) WhereEq(field string, target any) ([]*MessageRead, error) {
	query := bson.M{field: bson.M{"$all": target}}
	msgs, err := Where(messages, query, MessageRead{})

	return msgs, err
}

func (m Message) FindId(_id primitive.ObjectID) (any, error) {
	var result MessageRead
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

func (m Message) Create(data MessageWrite) error {
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	data.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err := messages.InsertOne(Ctx, data)

	return err
}

func (m Message) Update(data map[string]any, id interface{}) error {
	data["updated_at"] = primitive.NewDateTimeFromTime(time.Now())

	_, err := messages.UpdateByID(Ctx, id, data)

	return err
}
