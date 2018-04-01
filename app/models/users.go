package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId uint `gorm: "primary_key"`
	Name   string
	Mail   string
}
