package models

import (
	"github.com/fchieff/book-store-management-go-mysql.git/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
