package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

const (
	MongoDBName = "postbar"
)

var (
	MongoClient *mongo.Client
	MysqlDB     *gorm.DB
	MysqlUrl    string
)
