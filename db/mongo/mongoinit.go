package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"postbar/utils"
	"time"
)

var (
	client *mongo.Client
)

func init() {
	host, port, err := utils.GetMongoENV()
	if err != nil {
		log.Printf("error while loading environment: %v", err)
		// TODO: 将错误信息写入RabbitMQ队列，进而将错误消息一个一个插入mysql数据库
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+host+":"+port))
}
