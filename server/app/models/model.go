package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database
var Ctx context.Context

func Init(mc *mongo.Database, ctx context.Context) {
	DB = mc
	Ctx = ctx
}
