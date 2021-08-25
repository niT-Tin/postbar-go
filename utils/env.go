package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetMongoENV() (string, string, error) {
	err := godotenv.Load("../envs/mongo.env")
	if err != nil {
		log.Printf("error while loading environment: %v", err)
		return "", "", err
		// TODO: 将错误信息写入RabbitMQ队列，进而将错误消息一个一个插入mysql数据库
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if len(host) == 0 || len(port) == 0 {
		log.Printf("error while getting environment: %v", err)
		// TODO: 将错误信息写入RabbitMQ队列，进而将错误消息一个一个插入mysql数据库
		return "", "", err
	}
	return host, port, nil
}
