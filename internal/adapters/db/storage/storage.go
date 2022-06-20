package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
)

const (
	carTable       = "car"
	userTable      = "users"
	garageTable    = "garage"
	garageCarTable = "garage_car"
)

type Storage struct {
	Car
	User
	Garage
	logger *myLogger.Logger
}

func NewStorages(db *sqlx.DB, logger *myLogger.Logger) *Storage {

	carStorage := NewCarStorage(db, logger)
	userStorage := NewUserStorage(db, logger)
	garageStorage := NewGarageStorage(db, logger)

	return &Storage{
		Car:    carStorage,
		User:   userStorage,
		Garage: garageStorage,
	}
}
