package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/entity"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
)

type User interface {
	Create(dto dto.CreateUserDto) (int64, error)
	GetById(id int64) (*entity.User, error)
}

type userStorage struct {
	db     *sqlx.DB
	logger *myLogger.Logger
}

func NewUserStorage(db *sqlx.DB, logger *myLogger.Logger) *userStorage {
	return &userStorage{db: db, logger: logger}
}

func (st *userStorage) Create(dto dto.CreateUserDto) (int64, error) {

	var id int64

	sql := fmt.Sprint("INSERT INTO users (name,salary) values ($1,$2) returning user_id")
	st.logger.PrintWithInf(fmt.Sprintf("executing %s", sql))
	row := st.db.QueryRow(sql, dto.Name, dto.Salary)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

func (st *userStorage) GetById(id int64) (*entity.User, error) {

	var user entity.User

	sql := fmt.Sprint("SELECT * FROM users WHERE user_id=$1")
	st.logger.PrintWithInf(fmt.Sprintf("executing %s", sql))

	row := st.db.QueryRow(sql, id)

	if err := row.Scan(&user.UserID, &user.Name, &user.Salary); err != nil {
		return &user, err
	}

	return &user, nil

}
