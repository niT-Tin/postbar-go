package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"postbar/datamodels"
	"postbar/err"
)

type IContent interface {
	Create(*datamodels.Content) (*mongo.InsertOneResult, error)
	Update(*datamodels.Content) (*mongo.UpdateResult, error)
	Delete(*datamodels.Content) (*mongo.DeleteResult, error)
	GetOneById(int64) (*datamodels.Content, error)
	GetAll() ([]datamodels.Content, error)
	CheckRight() bool
}

type ContentRepository struct {
	db             string
	collectionName string
	mongodb        *mongo.Client
	collection     *mongo.Collection
}

func reciteErrorInRepo(errI *error) bool {
	var err2 error
	if errI != nil {
		err2 = errors.New((*errI).Error() + " on repo")
		return err.ReciteErr(&err2)
	}
	return false
}

func (c *ContentRepository) CheckRight() bool {
	if len(c.collectionName) == 0 || len(c.db) == 0 || c.mongodb == nil { //简单检查参数是否正确
		err2 := errors.New("new content failed")
		err.ReciteErr(&err2) //错误则将错误信息写入数据库
		return false
	}
	return true
}

func NewContentRepository(dbs, cl string, m *mongo.Client) IContent {
	c := &ContentRepository{
		db:             dbs,
		collectionName: cl,
		mongodb:        m,
	}
	if !c.CheckRight() {
		return nil
	}
	c.collection = c.mongodb.Database(c.db).Collection(c.collectionName)
	return c
}

// Create 插入单个内容文档
func (c *ContentRepository) Create(cnt *datamodels.Content) (*mongo.InsertOneResult, error) {
	res, err2 := c.collection.InsertOne(context.TODO(), cnt)
	if err2 != nil {
		reciteErrorInRepo(&err2)
	}
	return res, err2
}

// Update 更新单个内容文档
func (c *ContentRepository) Update(cnt *datamodels.Content) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"content_id": cnt.ContentId,
	}
	res, err2 := c.collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", cnt}})
	if err2 != nil {
		reciteErrorInRepo(&err2)
	}
	return res, err2
}

// Delete 删除单个内容文档
func (c *ContentRepository) Delete(cnt *datamodels.Content) (*mongo.DeleteResult, error) {
	filter := bson.D{
		{"content_id", cnt.ContentId},
	}
	res, err2 := c.collection.DeleteOne(context.TODO(), filter)
	if err2 != nil {
		reciteErrorInRepo(&err2)
	}
	return res, err2
}

// GetOneById 获取单个内容文档
func (c *ContentRepository) GetOneById(commentId int64) (*datamodels.Content, error) {
	one := c.collection.FindOne(context.TODO(), bson.M{"comment_id": commentId})
	if one == nil {
		err2 := errors.New("findOneById error res is nil")
		err.ReciteErr(&err2)
	}
	d := &datamodels.Content{}
	err2 := one.Decode(d)
	if err.ReciteErr(&err2) {
		return nil, err2
	}
	return d, nil
}

// GetAll 获取所有内容文档
func (c *ContentRepository) GetAll() ([]datamodels.Content, error) {
	var rtnRes []datamodels.Content
	result, err2 := c.collection.Find(context.TODO(), bson.D{})
	if err2 != nil {
		reciteErrorInRepo(&err2)
	}
	if err3 := result.All(context.TODO(), &rtnRes); err3 != nil {
		reciteErrorInRepo(&err3)
		return nil, err3
	}
	return rtnRes, nil
}
