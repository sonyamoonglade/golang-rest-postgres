package entity

type User struct {
	UserID int64   `json:"-" db:"user_id"`
	Name   string  `json:"name" db:"name"`
	Worth  float64 `json:"worth" db:"worth"`
}
