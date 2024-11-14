package model

import "time"

type Task struct {
	Id          int       `gorm:"type:int;primary_key"`
	Title       string    `gorm:"type:varchar(255);"`
	Description string    `gorm:"type:varchar(255)"`
	DueDate     time.Time `gorm:"type:date"`
}
