package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null;default:null"`
	Content string `json:"content" gorm:"not null;default:null"`
}
