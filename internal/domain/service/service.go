package service

import "github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"

type Service struct {
	Car
	Garage
	User
}

func NewServices(storages *storage.Storage) *Service {

	carService := NewCarService(storages.Car)
	userService := NewUserService(storages.User)
	garageService := NewGarageService(storages.Garage)
	return &Service{
		Car:    carService,
		User:   userService,
		Garage: garageService,
	}
}
