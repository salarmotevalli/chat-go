package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database
var Ctx context.Context

type Query interface {
	All() ([]interface{}, error)
	WhereEq(field string, target any) ([]interface{}, error)
	FindId(_id primitive.ObjectID) (any, error)
}

func Init(mc *mongo.Database, ctx context.Context) {
	DB = mc
	Ctx = ctx
}
