package storage

import "github.com/jmoiron/sqlx"

type Storage struct {
	Car
}

func NewStorages(db *sqlx.DB) *Storage {

	carStorage := NewCarStorage(db)

	return &Storage{
		Car: carStorage,
	}
}
