package datamodels

import "gorm.io/gorm"

type Router struct {
	gorm.Model
	ID           int64  `json:"id" gorm:"id"`
	RouterString string `json:"router_string" gorm:"router_string"`
	Description  string `json:"description" gorm:"description"`
}
