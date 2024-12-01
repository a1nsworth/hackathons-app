package repositories

import (
	"hackathons-app/internal/db"
)

type baseRepository struct {
	connection *db.GormDatabase
}

func (b baseRepository) Connection() *db.GormDatabase {
	return b.connection
}

func newBaseRepository(gormDb *db.GormDatabase) baseRepository {
	return baseRepository{connection: gormDb}
}
