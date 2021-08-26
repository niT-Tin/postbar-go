package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"postbar/datamodels"
	"postbar/db"
	"postbar/err"
)

type ISinglePostRepository interface {
	Create(post *datamodels.SinglePost) error
	Delete(int64) error
	GetCommentsInPostByPostId(int64) ([]datamodels.Comment, error)
	GetOneById(postId int64) (*datamodels.SinglePost, error)
	CrudSingleComment() IComment
	Update(post *datamodels.SinglePost) error
	CheckRight() bool
}

type SinglePostRepository struct {
	db             string
	collectionName string
	mongodb        *mongo.Client
	collection     *mongo.Collection
}

func (s *SinglePostRepository) CheckRight() bool {
	if len(s.collectionName) == 0 || len(s.db) == 0 || s.mongodb == nil { //简单检查参数是否正确
		err2 := errors.New("new content failed")
		err.ReciteErr(&err2) //错误则将错误信息写入数据库
		return false
	}
	return true
}

func NewSinglePostRepository(db string, collectionName string, s *mongo.Client) ISinglePostRepository {
	s1 := &SinglePostRepository{
		db:             db,
		collectionName: collectionName,
		mongodb:        s,
	}
	if !s1.CheckRight() {
		return nil
	}
	s1.collection = s1.mongodb.Database(s1.db).Collection(s1.collectionName)
	return s1
}

func (s *SinglePostRepository) Create(post *datamodels.SinglePost) error {
	// 实现id自增
	db.SinglePostIdInc += 1
	post.PosterId = db.SinglePostIdInc
	_, err2 := s.collection.InsertOne(context.TODO(), post)
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}

func (s *SinglePostRepository) Delete(postId int64) error {
	_, err2 := s.collection.DeleteOne(context.TODO(), bson.D{{"single_post_id", postId}})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}

func (s *SinglePostRepository) GetCommentsInPostByPostId(postId int64) ([]datamodels.Comment, error) {
	id, err2 := s.GetOneById(postId)
	return id.Comments, err2
}

func (s *SinglePostRepository) CrudSingleComment() IComment {
	return NewCommentRepository(db.MongoDBName, "comment", s.mongodb)
}

func (s *SinglePostRepository) GetOneById(postId int64) (*datamodels.SinglePost, error) {
	one := s.collection.FindOne(context.TODO(), bson.D{{"single_post_id", postId}})
	if e := one.Err(); e != nil {
		reciteErrorInRepo(&e)
		return nil, e
	}
	p := &datamodels.SinglePost{}
	err2 := one.Decode(p)
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return nil, err2
	}
	return p, nil
}

func (s *SinglePostRepository) Update(post *datamodels.SinglePost) error {
	_, err2 := s.collection.UpdateOne(context.TODO(), bson.D{{"single_post_id", post.SinglePostId}}, bson.D{{"$set", post}})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}
