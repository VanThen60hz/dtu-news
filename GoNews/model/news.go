package model

import "gorm.io/gorm"

type News struct {
	ID      int    `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Title   string `gorm:"type:varchar(50);not null" json:"title"`
	Summary string `gorm:"type:varchar(100)" json:"summary"`
	Content string `gorm:"type:varchar(255);not null" json:"content"`
	*gorm.Model
}
