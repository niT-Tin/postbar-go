package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"postbar/datamodels"
	"postbar/err"
)

type IComment interface {
	Insert(*datamodels.Comment) (*mongo.InsertOneResult, error)
	//GetAllInSinglePost(singlePostId int64) ([]datamodels.Comment, error)
	GetOneById(int64) (*datamodels.Comment, error)
	Update(*datamodels.Comment) (*mongo.UpdateResult, error)
	Delete(*datamodels.Comment) error
	GetContentInCommentById(int64) (*datamodels.Content, error)
	CheckRight() bool
}

type CommentRepository struct {
	db             string
	collectionName string
	mongodb        *mongo.Client
	collection     *mongo.Collection
}

func (m *CommentRepository) CheckRight() bool {
	if len(m.collectionName) == 0 || len(m.db) == 0 || m.mongodb == nil { //简单检查参数是否正确
		err2 := errors.New("new content failed")
		err.Reciteerr(&err2) //错误则将错误信息写入数据库
		return false
	}
	return true
}

func NewCommentRepository(db, collectionName string, c *mongo.Client) IComment {
	t := &CommentRepository{
		db:             db,
		collectionName: collectionName,
		mongodb:        c,
	}
	if !t.CheckRight() {
		return nil
	}
	t.collection = t.mongodb.Database(t.db).Collection(t.collectionName)
	return t
}

// Insert 插入单条评论
func (m *CommentRepository) Insert(comment *datamodels.Comment) (*mongo.InsertOneResult, error) {
	one, err := m.collection.InsertOne(context.TODO(), comment)
	if err != nil {
		reciteErrorInRepo(&err)
		return nil, err
	}
	return one, nil
}

//func (m *CommentRepository) GetAllInSinglePost(singlePostId int64) ([]datamodels.Comment, error) {
//
//}

func (m *CommentRepository) GetOneById(commentId int64) (*datamodels.Comment, error) {
	one := m.collection.FindOne(context.TODO(), bson.M{"comment_id": commentId})
	if one == nil {
		e := errors.New("查找失败")
		reciteErrorInRepo(&e)
		return nil, e
	}
	c := &datamodels.Comment{}
	err := one.Decode(c)
	if err != nil {
		reciteErrorInRepo(&err)
		return nil, err
	}
	return c, nil
}

func (m *CommentRepository) Update(newComment *datamodels.Comment) (*mongo.UpdateResult, error) {
	id, err := m.collection.UpdateOne(context.TODO(), bson.D{{"comment_id", newComment.CommentId}}, bson.D{{"$set", newComment}})

	if err != nil {
		reciteErrorInRepo(&err)
		return nil, err
	}
	return id, nil
}

func (m *CommentRepository) Delete(cmt *datamodels.Comment) error {
	_, err := m.collection.DeleteOne(context.TODO(), bson.D{{"comment_id", cmt.CommentId}})
	if err != nil {
		reciteErrorInRepo(&err)
		return err
	}
	return nil
}

func (m *CommentRepository) GetContentInCommentById(commentid int64) (*datamodels.Content, error) {
	id, err2 := m.GetOneById(commentid)
	if err2 != nil {
		reciteErrorInRepo(&err2)
	}
	return &id.Contents, nil
}
