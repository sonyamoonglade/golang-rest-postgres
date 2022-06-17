package service

import (
	"github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/entity"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
)

type Car interface {
	CreateCar(dto dto.CreateCarDto) (uint8, error)
	DeleteCar(id uint8) (uint8, error)
	GetCar(id uint8) (*entity.Car, error)
	UpdateCar(dto dto.UpdateCarDto) (*entity.Car, error)
}

type CarService struct {
	storage storage.Car
}

func NewCarService(storage storage.Car) *CarService {
	return &CarService{storage: storage}
}

func (s *CarService) CreateCar(dto dto.CreateCarDto) (uint8, error) {
	return 0, nil
}
func (s *CarService) UpdateCar(dto dto.UpdateCarDto) (*entity.Car, error) {
	return nil, nil
}
func (s *CarService) DeleteCar(id uint8) (uint8, error) {
	return 0, nil
}
func (s *CarService) GetCar(id uint8) (*entity.Car, error) {
	return nil, nil
}
