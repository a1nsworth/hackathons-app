package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName      string
	SecondName     string
	Email          string
	TelegramID     string
	HashedPassword string
	Hackathons     []Hackathon `gorm:"many2many:hackathons_users;"`
}
