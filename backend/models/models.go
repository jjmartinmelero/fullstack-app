package models

import (
	"backend-fullstack-app/database"
	"time"
)

type Category struct {
	Id   uint   `json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Slug string `gorm:"type:varchar(100)" json:"slug"`
}

type Categories []Category

type Recipe struct {
	Id          uint      `json:"id"`
	CategoryId  uint      `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryId" json:"category"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Slug        string    `gorm:"type:varchar(100)" json:"slug"`
	Time        string    `gorm:"type:varchar(100);not null" json:"time"`
	Photo       string    `gorm:"type:varchar(100);not null" json:"photo"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Recipes []Recipe

type Contact struct {
	Id      uint      `json:"id"`
	Name    string    `gorm:"type:varchar(100);not null" json:"name"`
	Email   string    `gorm:"type:varchar(100);not null" json:"email"`
	Phone   string    `gorm:"type:varchar(100);not null" json:"phone"`
	Message string    `gorm:"type:text;not null" json:"message"`
	Date    time.Time `json:"date"`
}

type Contacts []Contact

func Migration() {
	database.DB.AutoMigrate(&Recipe{}, &Category{}, &Contact{})
}
