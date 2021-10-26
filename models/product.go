package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50); not null" json:"name"`
	Type   string `gorm:"type:varchar(10);" json:"type"`
	Price  int    `json:"price"`
	Stock  int    `gorm:"type:int(10);" json:"stock"`
	Images string `json:"images"`
}
