package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
)

type Car interface {
	Create(dto dto.CreateCarDto) (uint8, error)
}

type CarStorage struct {
	db *sqlx.DB
}

func NewCarStorage(db *sqlx.DB) *CarStorage {
	return &CarStorage{db: db}
}

func (s *CarStorage) Create(dto dto.CreateCarDto) (uint8, error) {
	var id uint8

	sql := fmt.Sprint("INSERT INTO car (garage_id,user_id,model,year,price) VALUES ($1,$2,$3,$4,$5)")
	row := s.db.QueryRow(sql, dto.GarageID, dto.UserID, dto.Model, dto.Year, dto.Price)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
