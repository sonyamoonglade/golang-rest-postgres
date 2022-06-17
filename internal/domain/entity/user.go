package entity

type User struct {
	UserID   uint8   `json:"-" db:"user_id"`
	GarageID uint8   `json:"garageId" db:"garage_id"`
	Name     string  `json:"name" db:"name"`
	Salary   float64 `json:"salary" db:"salary"`
}
