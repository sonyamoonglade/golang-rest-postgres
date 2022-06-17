package entity

type User struct {
	UserID   uint8   `util:"-" db:"user_id"`
	GarageID uint8   `util:"garageId" db:"garage_id"`
	Name     string  `util:"name" db:"name"`
	Salary   float64 `util:"salary" db:"salary"`
}
