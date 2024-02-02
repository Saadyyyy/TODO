package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"index"`
	Title  string `json:"title" gorm:"type:varchar(255)"`
	Body   string `json:"body" gorm:"type:varchar(255)"`
	Status bool   `json:"status"`
	IsRead bool   `json:"is_read"`
}
