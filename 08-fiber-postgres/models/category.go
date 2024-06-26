package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"foreignKey:CategoryID"`
}
