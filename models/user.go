package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(50); not null, unique" json:"username"`
	Fullname    string `gorm:"not null" json:"fullname"`
	Email       string `gorm:"type:varchar(75);unique_index" json:"email"`
	Password    string `gorm:"not null" json:",omitempty"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
}
