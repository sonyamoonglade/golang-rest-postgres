package service

import "github.com/sonyamoonglade/golang-rest-postgres/pkg/repository"

type Service struct {
	Car
	Repositories *repository.Repository
}

type Car interface {
	createCar()
	getCar()
}

func CreateService(repository *repository.Repository) *Service {
	s := Service{Repositories: repository}
	return &s
}
