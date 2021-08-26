package repositories

import (
	"errors"
	"gorm.io/gorm"
	"postbar/datamodels"
)

type IError interface {
	Insert(error2 *datamodels.Error)
	SelectAll() []datamodels.Error
}

type ErrorS struct {
	db *gorm.DB
}

func NewErrorS(db *gorm.DB) IError {
	if db == nil {
		var e = errors.New("mysql数据库初始化错误")
		reciteErrorInRepo(&e)
		return nil
	}
	return &ErrorS{db: db}
}

func (e *ErrorS) Insert(error2 *datamodels.Error) {
	e.db.Create(error2)
}

func (e *ErrorS) SelectAll() []datamodels.Error {
	var t []datamodels.Error
	find := e.db.Find(&t)
	if er := find.Error; er != nil {
		reciteErrorInRepo(&er)
		return nil
	}
	return t
}
