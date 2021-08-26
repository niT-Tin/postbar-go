package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

const (
	MongoDBName              = "postbar"
	CommentCollectionName    = "comment"
	ContentCollectionName    = "content"
	SinglePostCollectionName = "single_post"
	PostBarCollectionName    = "post"
	UserCollectionName       = "user"
)

var (
	MongoClient *mongo.Client
	MysqlDB     *gorm.DB
	MysqlUrl    string
)
