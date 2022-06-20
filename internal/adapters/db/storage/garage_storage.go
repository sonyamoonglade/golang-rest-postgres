package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
)

type Garage interface {
	Create(dto dto.CreateGarageDto) (int64, error)
}

type garageStorage struct {
	db     *sqlx.DB
	logger *myLogger.Logger
}

func NewGarageStorage(db *sqlx.DB, logger *myLogger.Logger) *garageStorage {
	return &garageStorage{db: db, logger: logger}
}

func (s *garageStorage) Create(dto dto.CreateGarageDto) (int64, error) {

	var garageId int64

	sql := fmt.Sprint("INSERT INTO garage (user_id, name, capacity,space_left) values ($1,$2,$3,$4) returning garage_id")
	s.logger.PrintWithInf(fmt.Sprintf("executing %s", sql))
	row := s.db.QueryRow(sql, dto.UserID, dto.Name, dto.Capacity, dto.Capacity)

	if err := row.Scan(&garageId); err != nil {
		return 0, err
	}

	return garageId, nil
}
