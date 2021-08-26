package datamodels

import "gorm.io/gorm"

type Error struct {
	gorm.Model
	Message string `json:"message" gorm:"<-:false"`
	Err     error  `json:"error" gorm:"<-:false"`
}
