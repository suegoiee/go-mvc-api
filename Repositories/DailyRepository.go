package Repositories

import (
	"context"
	"fmt"
	"strconv"

	"ginPrac/Models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type DailyRepository struct {
	collection *mongo.Collection
}

func NewDailyRepository(collection *mongo.Collection) *DailyRepository {
	return &DailyRepository{
		collection: collection,
	}
}

func (ur *DailyRepository) GetDaily(code string) (*Models.Daily, error) {
	var user Models.Daily
	intCode, convertErr := strconv.Atoi(code)
	filter := bson.M{"Code": intCode}

	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil && convertErr != nil {
		return nil, fmt.Errorf("failed to get user: %v the code is %s", err, code)
	}

	return &user, nil
}

// func (ur *UserRepository) CreateUser(user model.User) (*model.User, error) {
// 	_, err := ur.collection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create user: %v", err)
// 	}

// 	return &user, nil
// }
