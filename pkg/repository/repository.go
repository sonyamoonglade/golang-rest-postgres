package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/entities"
)

type Repository struct {
	Car
	User
}

type Car interface {
}

type User interface {
	CreateUser(user entities.User) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
