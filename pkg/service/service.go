package service

import (
	"github.com/sonyamoonglade/golang-rest-postgres/entities"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/repository"
)

type Service struct {
	Car
	User
}

type Car interface {
	CreateCar()
	GetCar()
}

type User interface {
	CreateUser(u entities.User) (int, error)
}

func CreateService(r *repository.Repository) *Service {
	s := Service{
		User: NewUserService(r.User),
	}
	return &s
}
