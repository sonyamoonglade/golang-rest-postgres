package service

import (
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
)

type Garage interface {
	Create(dto dto.CreateGarageDto) (int64, error)
}

type garageService struct {
	storage storage.Garage
}

func NewGarageService(storage storage.Garage) *garageService {
	return &garageService{storage: storage}
}

func (s *garageService) Create(dto dto.CreateGarageDto) (int64, error) {
	fmt.Println("here")
	garageId, err := s.storage.Create(dto)
	if err != nil {
		return 0, err
	}

	return garageId, nil

}
