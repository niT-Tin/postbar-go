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

type IUserRepository interface {
	Insert(user *datamodels.User) error
	Delete(userid int64) error
	Update(user *datamodels.User) error
	GetOneById(int64) *datamodels.User
	GetAllUsers() []datamodels.User
	CheckRight() bool
	GetUserByAccount(string) (*datamodels.User, error)
	PostCrud() ISinglePostRepository
	CommentCrud() IComment
	PostBarCrud() IPostBar
}

func (u *UserRepository) CheckRight() bool {
	if len(u.collectionName) == 0 || len(u.db) == 0 || u.mongodb == nil { //简单检查参数是否正确
		err2 := errors.New("new content failed")
		err.ReciteErr(&err2) //错误则将错误信息写入数据库
		return false
	}
	return true
}

type UserRepository struct {
	db             string
	collectionName string
	mongodb        *mongo.Client
	collection     *mongo.Collection
}

func NewUserRepository(db, collectionName string, u1 *mongo.Client) IUserRepository {
	tu := &UserRepository{
		db:             db,
		collectionName: collectionName,
		mongodb:        u1,
	}
	if !tu.CheckRight() {
		return nil
	}
	tu.collection = tu.mongodb.Database(tu.db).Collection(tu.collectionName)
	return tu
}

func (u *UserRepository) Insert(user *datamodels.User) error {
	// 实现id自增
	db.UserIdInc += 1
	user.Userid = db.UserIdInc
	_, err2 := u.collection.InsertOne(context.TODO(), user)
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}

func (u *UserRepository) Delete(userid int64) error {
	_, err2 := u.collection.DeleteOne(context.TODO(), bson.D{{"userid", userid}})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}

func (u *UserRepository) Update(user *datamodels.User) error {
	_, err2 := u.collection.UpdateOne(context.TODO(), bson.D{{"userid", user.Userid}}, bson.D{{"$set", user}})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}

func (u *UserRepository) GetOneById(userid int64) *datamodels.User {
	one := u.collection.FindOne(context.TODO(), bson.D{{"userid", userid}})
	if e := one.Err(); e != nil {
		reciteErrorInRepo(&e)
		return nil
	}
	t := &datamodels.User{}
	err2 := one.Decode(t)
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return nil
	}
	return t
}

func (u *UserRepository) GetAllUsers() []datamodels.User {
	find, err2 := u.collection.Find(context.TODO(), bson.D{})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return nil
	}
	var us []datamodels.User
	if e := find.All(context.TODO(), &us); e != nil {
		reciteErrorInRepo(&e)
		return nil
	}
	return us
}

func (u *UserRepository) GetUserByAccount(account string) (*datamodels.User, error) {
	one := u.collection.FindOne(context.TODO(), bson.D{{"account", account}})
	if e := one.Err(); e != nil {
		reciteErrorInRepo(&e)
		return nil, e
	}
	t := &datamodels.User{}
	err2 := one.Decode(t)
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return nil, err2
	}
	return t, nil
}

func (u *UserRepository) PostCrud() ISinglePostRepository {
	return NewSinglePostRepository(db.MongoDBName, db.SinglePostCollectionName, u.mongodb)
}

func (u *UserRepository) CommentCrud() IComment {
	return NewCommentRepository(db.MongoDBName, db.CommentCollectionName, u.mongodb)
}

func (u *UserRepository) PostBarCrud() IPostBar {
	return NewPostBarRepository(db.MongoDBName, db.PostBarCollectionName, u.mongodb)
}
