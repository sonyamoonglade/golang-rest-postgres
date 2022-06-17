package storage

import "github.com/jmoiron/sqlx"

type Car interface {
}

type CarStorage struct {
	db *sqlx.DB
}

func NewCarStorage(db *sqlx.DB) *CarStorage {
	return &CarStorage{db: db}
}
