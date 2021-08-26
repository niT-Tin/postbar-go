package repositories

import (
	"gorm.io/gorm"
	"postbar/datamodels"
)

type IIdsRepository interface {
	Insert(idVar *datamodels.IdVar) bool
	Update(idVar *datamodels.IdVar) bool
	Delete(idVar *datamodels.IdVar) bool
}

type IdsRepository struct {
	db *gorm.DB
}

func NewIdsRepository(db *gorm.DB) IIdsRepository {
	return &IdsRepository{db: db}
}

func (i *IdsRepository) Insert(idVar *datamodels.IdVar) bool {
	create := i.db.Create(idVar)
	if create.RowsAffected != 0 {
		return true
	}
	return false
}

func (i *IdsRepository) Update(idVar *datamodels.IdVar) bool {
	updates := i.db.Model(idVar).Updates(idVar)
	if updates.RowsAffected != 0 {
		return true
	}
	return false
}

func (i *IdsRepository) Delete(idVar *datamodels.IdVar) bool {
	tx := i.db.Model(idVar).Delete(idVar)
	if tx.RowsAffected != 0 {
		return true
	}
	return false
}
