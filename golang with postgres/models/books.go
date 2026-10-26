package models

import "gorm.io/gorm"

type books struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Punlisher string `json:"publisher"`
}
