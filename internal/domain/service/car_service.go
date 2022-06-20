package service

import (
	"context"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/entity"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/middleware"
)

type Car interface {
	CreateCar(dto dto.CreateCarDto) (int64, error)
	DeleteCar(id int64) (int64, error)
	GetCar(id int64) (*entity.Car, error)
	UpdateCar(dto dto.UpdateCarDto) (*entity.Car, error)
	BuyCar(dto dto.BuyCarDto, ctx context.Context) (bool, error)
}

type carService struct {
	storage storage.Car
}

func NewCarService(storage storage.Car) *carService {
	return &carService{storage: storage}
}

func (s *carService) CreateCar(dto dto.CreateCarDto) (int64, error) {

	carId, err := s.storage.Create(dto)
	if err != nil {
		return 0, err
	}

	return carId, nil

}
func (s *carService) UpdateCar(dto dto.UpdateCarDto) (*entity.Car, error) {
	return nil, nil
}
func (s *carService) DeleteCar(id int64) (int64, error) {
	return 0, nil
}
func (s *carService) GetCar(id int64) (*entity.Car, error) {
	return nil, nil
}
func (s *carService) BuyCar(dto dto.BuyCarDto, ctx context.Context) (bool, error) {

	userId := ctx.Value(middleware.UserId).(int64)
	dto.UserID = userId

	ok, err := s.storage.Buy(dto)
	if err != nil {
		return ok, err
	}

	return ok, nil
}
