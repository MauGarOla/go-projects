package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}
