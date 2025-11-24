package models

import "backend-fullstack-app/database"

type Category struct {
	Id   uint   `json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Slug string `gorm:"type:varchar(100)" json:"slug"`
}

type Categories []Category

func Migration() {
	database.DB.AutoMigrate(&Category{})
}
