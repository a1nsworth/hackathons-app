package models

import (
	"time"

	"gorm.io/gorm"
)

type Hackathon struct {
	gorm.Model
	Name        string
	Description string
	DateBegin   time.Time
	DateEnd     time.Time
	Users       []User `gorm:"many2many:hackathons_users;"`
}
