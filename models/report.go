package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Approved    bool   `json:"approved"`
}
