package main

import (
	"context"
	"ginPrac/Controllers"
	"ginPrac/Repositories"
	"ginPrac/Services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)

func main() { // 建立MongoDB連接
	credential := options.Credential{
		Username: "",
		Password: "",
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 指定要使用的數據庫和集合
	db := client.Database("stock")
	collection := db.Collection("news")

	dailyRepository := Repositories.NewDailyRepository(collection)
	dailyService := Services.NewDailyService(dailyRepository)
	dailyController := Controllers.NewDailyController(dailyService)

	route := gin.Default()

	route.GET("/today/:code", dailyController.PriceToday)
	route.Run()
}
