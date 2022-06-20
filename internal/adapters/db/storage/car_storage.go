package storage

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/app_errors"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
)

type Car interface {
	Create(dto dto.CreateCarDto) (int64, error)
	Buy(dto dto.BuyCarDto) (bool, error)
}

type carStorage struct {
	db     *sqlx.DB
	logger *myLogger.Logger
}

func NewCarStorage(db *sqlx.DB, logger *myLogger.Logger) *carStorage {
	return &carStorage{db: db, logger: logger}
}

func (s *carStorage) Create(dto dto.CreateCarDto) (int64, error) {

	var carId int64

	sql := fmt.Sprintf("INSERT INTO %s (model,year,price) VALUES ($1,$2,$3) returning car_id", carTable)
	s.logger.PrintWithInf(fmt.Sprintf("executing %s", sql))

	row := s.db.QueryRow(sql, dto.Model, dto.Year, dto.Price)
	if err := row.Scan(&carId); err != nil {
		return 0, err
	}
	return carId, nil
}
func (s *carStorage) Buy(dto dto.BuyCarDto) (bool, error) {

	var carPrice float64
	var spaceLeft int64
	var worth float64

	tx, err := s.db.Beginx()
	if err != nil {
		tx.Rollback()
		s.logger.PrintWithErr(fmt.Sprintf("tx has failed. %s", err.Error()))
		return false, err
	}
	s.logger.PrintWithInf("transaction has began")

	{

		q := fmt.Sprintf("INSERT INTO %s (car_id,garage_id,user_id) VALUES ($1,$2,$3)", garageCarTable)
		_, err = tx.Exec(q, dto.CarID, dto.GarageID, dto.UserID)
		s.logger.PrintWithInf(fmt.Sprintf("executing %s", q))
		if err != nil {
			tx.Rollback()
			s.logger.PrintWithErr(fmt.Sprintf("tx has failed. %s", err.Error()))
			return false, nil
		}

		q1 := fmt.Sprintf("SELECT price from %s WHERE car_id = $1", carTable)
		row := tx.QueryRow(q1, dto.CarID)
		s.logger.PrintWithInf(fmt.Sprintf("executing %s", q))

		if err := row.Scan(&carPrice); err != nil {
			tx.Rollback()
			s.logger.PrintWithErr(fmt.Sprintf("tx has failed. %s", err.Error()))
			return false, err
		}

		q2 := fmt.Sprintf("UPDATE %s SET worth = CASE WHEN worth > $1 THEN worth - $1 ELSE worth END WHERE user_id = $2 RETURNING worth", userTable)
		row = tx.QueryRow(q2, carPrice, dto.UserID)
		s.logger.PrintWithInf(fmt.Sprintf("executing %s", q2))
		if err := row.Scan(&worth); err != nil {
			tx.Rollback()
			return false, err
		}

		if worth < carPrice {
			s.logger.PrintWithErr("tx is rolling back. Worth is not enough")
			msg := fmt.Sprintf("worth is not enough to buy car with id %d", dto.CarID)
			return false, app_errors.NewApiError(msg, errors.New(msg))
		}

		q3 := fmt.Sprintf("UPDATE %s SET space_left = CASE WHEN space_left > 0 THEN space_left - 1 ELSE space_left END RETURNING space_left", garageTable)
		row = tx.QueryRow(q3)
		s.logger.PrintWithInf(fmt.Sprintf("executing %s", q3))
		if err := row.Scan(&spaceLeft); err != nil {
			tx.Rollback()
			s.logger.PrintWithErr(fmt.Sprintf("tx has failed. %s", err.Error()))
			return false, err
		}

		if spaceLeft == 0 {
			tx.Rollback()
			s.logger.PrintWithErr("tx is rolling back. space_left = 0")
			msg := "can't fit car in garage with space_left = 0"
			return false, app_errors.NewApiError(msg, errors.New(msg))
		}

	}

	tx.Commit()
	s.logger.PrintWithInf("transaction has committed")
	if err != nil {
		return false, err
	}

	return true, nil

}
