package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func Where[T any](collection *mongo.Collection, query any, t T) ([]*T, error) {

	var result []*T
	// Run query
	cur, err := collection.Find(Ctx, query)
	if err != nil {
		return nil, err
	}

	// Collect data
	for cur.Next(Ctx) {
		var item T

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

	if result == nil {
		result = []*T{}
	}

	return result, nil
}
