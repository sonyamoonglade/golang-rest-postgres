package entity

type User struct {
	UserID int64  `json:"-" db:"user_id"`
	Name   string `json:"name" db:"name"`
	Salary int64  `json:"salary" db:"salary"`
}
