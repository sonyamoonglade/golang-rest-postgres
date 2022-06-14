package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/entities"
)

const (
	usersTable  = "users"
	garageTable = "garage"
	carTable    = "car"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(u entities.User) (int, error) {

	var userId uint8

	sql := fmt.Sprintf("insert into %s (name,salary) values ($1,$2) returning user_id", usersTable)
	row := r.db.QueryRow(sql, u.Name, u.Salary)

	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return int(userId), nil
}
