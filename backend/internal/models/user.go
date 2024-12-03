package models

import "gorm.io/gorm"

type Role int

const (
	Admin Role = iota
	Base
)

type User struct {
	gorm.Model
	FirstName      string
	SecondName     string
	Email          string
	TelegramID     string
	HashedPassword string
	Role           Role        `gorm:"default:1"`
	Hackathons     []Hackathon `gorm:"many2many:hackathons_users;"`
}

func a() Role {
	return Role(1)
}
