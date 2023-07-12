package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func Where(collection *mongo.Collection, query any) ([]any, error) {

	var result []*any
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

	var interfaces []any
	for _, item := range result {
		interfaces = append(interfaces, item)
	}

	if interfaces == nil {
		interfaces = []any{}
	}

	return interfaces, nil
}
