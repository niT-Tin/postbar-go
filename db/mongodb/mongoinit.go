package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"postbar/db"
	"postbar/utils"
)

func GetClient() *mongo.Client {
	return db.MongoClient
}

func init() {
	host, port, err := utils.GetMongoENV()
	if err != nil {
		log.Printf("error while loading environment: %v", err)
		// TODO: 将错误信息写入RabbitMQ队列，进而将错误消息一个一个插入mysql数据库
	}
	db.MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+host+":"+port))
}
