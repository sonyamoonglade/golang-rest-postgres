package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
)

type Car interface {
	Create(dto dto.CreateCarDto) (uint8, error)
}

type carStorage struct {
	db     *sqlx.DB
	logger *myLogger.Logger
}

func NewCarStorage(db *sqlx.DB, logger *myLogger.Logger) *carStorage {
	return &carStorage{db: db, logger: logger}
}

func (s *carStorage) Create(dto dto.CreateCarDto) (uint8, error) {
	var id uint8
	sql := fmt.Sprint("INSERT INTO car (garage_id,user_id,model,year,price) VALUES ($1,$2,$3,$4,$5)")
	s.logger.PrintWithInf(fmt.Sprintf("executing %s ", sql))
	row := s.db.QueryRow(sql, dto.GarageID, dto.UserID, dto.Model, dto.Year, dto.Price)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
