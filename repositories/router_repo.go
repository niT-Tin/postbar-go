package repositories

import (
	"errors"
	"gorm.io/gorm"
	"postbar/datamodels"
)

type IRouter interface {
	Insert(router *datamodels.Router) bool
	Update(*datamodels.Router) bool
	Delete(string) bool
	SelectAll() []datamodels.Router
	SelectById(int64) datamodels.Router
}

type RouterRepository struct {
	db *gorm.DB
}

func NewRouterRepository(db *gorm.DB) IRouter {
	if db == nil {
		var e = errors.New("mysql数据库初始化错误")
		reciteErrorInRepo(&e)
		return nil
	}
	return &RouterRepository{db: db}
}

func (r *RouterRepository) Insert(router *datamodels.Router) bool {
	create := r.db.Create(router)
	if create.RowsAffected == 0 {
		return false
	}
	return true
}

func (r *RouterRepository) Update(router *datamodels.Router) bool {
	updates := r.db.Model(router).Where("id = ?", router.ID).Updates(router)
	if updates.RowsAffected == 0 {
		return false
	}
	return true
}

func (r *RouterRepository) Delete(router string) bool {
	tx := r.db.Model(&datamodels.Router{}).Where("router_string=?", router).Delete(datamodels.Router{})
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}

func (r *RouterRepository) SelectAll() []datamodels.Router {
	var t []datamodels.Router
	r.db.Find(&t)
	return t
}

func (r *RouterRepository) SelectById(routerId int64) datamodels.Router {
	var t datamodels.Router
	r.db.First(&t, routerId)
	return t
}
