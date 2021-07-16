package utils

import (
	"MongoGift/internal/status"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCon *mongo.Collection

func MongoClient() *status.Response {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return status.MongoDBErr
	}

	log.Println(client)
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return status.MongoDBErr
	}
	MongoCon = client.Database("gift").Collection("login")
	fmt.Println("Connected to MongoDB!")
	return nil
}
