package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"postbar/datamodels"
	"postbar/err"
)

type IPostBar interface {
	Create(bar *datamodels.PostBar) error
	Delete(int64) error
	GetOneById(int64) (*datamodels.PostBar, error)
	GetAllPostBar() ([]datamodels.PostBar, error)
	Update(*datamodels.PostBar) error
	CheckRight() bool
}

type PostBarRepository struct {
	db             string
	collectionName string
	mongodb        *mongo.Client
	collection     *mongo.Collection
}

func (p *PostBarRepository) CheckRight() bool {
	if len(p.collectionName) == 0 || len(p.db) == 0 || p.mongodb == nil { //简单检查参数是否正确
		err2 := errors.New("new content failed")
		err.Reciteerr(&err2) //错误则将错误信息写入数据库
		return false
	}
	return true
}

func NewPostBarRepository(db, collectionName string, p *mongo.Client) IPostBar {
	tp := &PostBarRepository{
		db:             db,
		collectionName: collectionName,
		mongodb:        p,
	}
	if !tp.CheckRight() {
		return nil
	}
	tp.collection = tp.mongodb.Database(tp.db).Collection(tp.collectionName)
	return tp
}

func (p *PostBarRepository) Create(bar *datamodels.PostBar) error {
	_, err2 := p.collection.InsertOne(context.TODO(), bar)
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}

func (p *PostBarRepository) Delete(barId int64) error {
	_, err2 := p.collection.DeleteOne(context.TODO(), bson.D{{"postbar_id", barId}})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return nil
}

func (p *PostBarRepository) GetOneById(barId int64) (*datamodels.PostBar, error) {
	one := p.collection.FindOne(context.TODO(), bson.M{"postbar_id": barId})
	if e := one.Err(); e != nil {
		reciteErrorInRepo(&e)
		return nil, e
	}
	t := &datamodels.PostBar{}
	err2 := one.Decode(t)
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return nil, err2
	}
	return t, nil
}

func (p *PostBarRepository) GetAllPostBar() ([]datamodels.PostBar, error) {
	find, err2 := p.collection.Find(context.TODO(), bson.D{})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return nil, err2
	}
	var t []datamodels.PostBar
	if e := find.All(context.TODO(), &t); e != nil {
		reciteErrorInRepo(&e)
		return nil, e
	}
	return t, nil
}

func (p *PostBarRepository) Update(bar *datamodels.PostBar) error {
	_, err2 := p.collection.UpdateOne(context.TODO(), bson.D{{"postbar_id", bar.PostBarId}}, bson.D{{"$set", bar}})
	if err2 != nil {
		reciteErrorInRepo(&err2)
		return err2
	}
	return err2
}
