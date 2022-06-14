package service

import (
	"github.com/sonyamoonglade/golang-rest-postgres/entities"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(u entities.User) (int, error) {

	userId, err := s.repo.CreateUser(u)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
