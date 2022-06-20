package entity

type Car struct {
	CarID    uint8   `json:"carId"  db:"car_id"`
	UserID   uint8   `json:"userId" db:"user_id"`
	GarageID uint8   `json:"garageID" db:"garage_id"`
	Model    string  `json:"model" db:"model"`
	Year     uint8   `json:"year" db:"year"`
	Price    float64 `json:"price" db:"price"`
}
