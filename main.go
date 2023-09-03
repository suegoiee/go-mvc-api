package main

import (
	"context"
	"fmt"
	"ginPrac/Controllers"
	"ginPrac/Repositories"
	"ginPrac/Services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)

type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func main() { // 建立MongoDB連接
	fmt.Println("2023/09/03")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed loading env file", err)
	}

	mongoConfig := MongoConfig{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
	}

	credential := options.Credential{
		Username: mongoConfig.Username,
		Password: mongoConfig.Password,
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoConfig.Host + ":" + mongoConfig.Port).SetAuth(credential))
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

	newsRepository := Repositories.NewNewsRepository(collection)
	newsService := Services.NewNewsService(newsRepository)
	newsController := Controllers.NewNewsController(newsService)

	route := gin.Default()

	route.GET("/news", newsController.GetNews)
	route.Run(":8001")
}
