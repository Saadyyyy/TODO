package models

import (
	"gorm.io/gorm"
)

type Tugas struct {
	gorm.Model
	Task        string `json:"task" gorm:"type:varchar(255)"`
	Level       string `json:"level" gorm:"type:varchar(255)"`
	Deadline    string `json:"deadline" gorm:"type:varchar(255)"`
	Description string `json:"des" gorm:"type:varchar(255)"`
	Status      bool   `json:"status"`
}
