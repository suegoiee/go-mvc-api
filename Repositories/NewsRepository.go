package Repositories

import (
	"context"
	"fmt"
	"time"

	"ginPrac/Models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewsRepository struct {
	collection *mongo.Collection
}

func NewNewsRepository(collection *mongo.Collection) *NewsRepository {
	return &NewsRepository{
		collection: collection,
	}
}

func (nr *NewsRepository) Get(page, pageSize int) ([]*Models.News, error) {
	var newsList []*Models.News
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Find()
	options.SetSort(bson.M{"time": -1})

	// Calculate the number of documents to skip
	skip := (page - 1) * pageSize

	// Set the limit and skip options
	options.SetLimit(int64(pageSize))
	options.SetSkip(int64(skip))

	cursor, err := nr.collection.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil, fmt.Errorf("failed to get news -> %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news Models.News
		if err := cursor.Decode(&news); err != nil {
			return nil, fmt.Errorf("failed to decode news")
		}
		newsList = append(newsList, &news)
	}

	if len(newsList) == 0 {
		return nil, fmt.Errorf("no news found")
	}

	return newsList, nil
}
