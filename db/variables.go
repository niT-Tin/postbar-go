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

	// 用于ID自增
	ContentIdInc    int64
	CommentIdInc    int64
	SinglePostIdInc int64
	PostBarIdInc    int64
	UserIdInc       int64
)
