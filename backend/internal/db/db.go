package db

import (
	"gorm.io/gorm"
)

type GormDatabase struct {
	db *gorm.DB
}

func NewGormDatabase(dialector gorm.Dialector, config *gorm.Config) (*GormDatabase, error) {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}
	return &GormDatabase{db: db}, nil
}

func (gd GormDatabase) GetDB() *gorm.DB {
	return gd.db
}
