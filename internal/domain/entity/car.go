package entity

type Car struct {
	CarID int64   `json:"carId"  db:"car_id"`
	Model string  `json:"model" db:"model"`
	Year  int64   `json:"year" db:"year"`
	Price float64 `json:"price" db:"price"`
}
