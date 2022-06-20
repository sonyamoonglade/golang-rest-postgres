package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
)

type Storage struct {
	Car
	User
	logger myLogger.Logger
}

func NewStorages(db *sqlx.DB, logger *myLogger.Logger) *Storage {

	carStorage := NewCarStorage(db, logger)
	userStorage := NewUserStorage(db, logger)

	return &Storage{
		Car:  carStorage,
		User: userStorage,
	}
}
