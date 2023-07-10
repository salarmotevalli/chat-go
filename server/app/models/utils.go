package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Where(collection *mongo.Collection, query bson.M) ([]interface{}, error) {

	var result []*interface{}
	// Run query
	cur, err := collection.Find(Ctx, query)
	if err != nil {
		return nil, err
	}

	// Collect data
	for cur.Next(Ctx) {
		var item interface{}

		err := cur.Decode(&item)
		if err != nil {
			return nil, err
		}

		result = append(result, &item)
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

	if interfaces == nil {
		interfaces = []interface{}{}
	}

	return interfaces, nil
}
