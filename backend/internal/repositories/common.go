package repositories

import (
	"hackathons-app/internal/db"
)

type baseRepository struct {
	connection *db.GormDatabase
}

func newBaseRepository(gormDb *db.GormDatabase) baseRepository {
	return baseRepository{connection: gormDb}
}
