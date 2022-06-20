package service

import (
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/app_errors"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/entity"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
)

type User interface {
	Create(dto dto.CreateUserDto) (int64, error)
	GetById(id int64) (*entity.User, error)
}

type userService struct {
	storage storage.User
}

func NewUserService(storage storage.User) *userService {
	return &userService{storage: storage}
}

func (s *userService) Create(dto dto.CreateUserDto) (int64, error) {

	userId, err := s.storage.Create(dto)
	if err != nil {
		return 0, err
	}

	return userId, nil

}
func (s *userService) GetById(id int64) (*entity.User, error) {

	user, err := s.storage.GetById(id)

	if user.UserID == 0 && err != nil {
		msg := fmt.Sprintf("User with id %d does not exist", id)
		return nil, app_errors.NewApiError(msg, err)
	}

	return user, nil

}
