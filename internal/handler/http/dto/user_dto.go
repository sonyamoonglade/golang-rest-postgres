package dto

type CreateUserDto struct {
	Name   string `json:"name"`
	Salary int64  `json:"salary"`
}
