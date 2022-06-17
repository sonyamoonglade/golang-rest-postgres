package dto

type CreateCarDto struct {
	GarageID uint8   `util:"garageID"`
	UserID   uint8   `util:"userID"`
	Model    string  `util:"model"`
	Year     uint8   `util:"year"`
	Price    float64 `util:"price"`
}

type UpdateCarDto struct {
	CarID    uint8   `util:"carID"`
	GarageID uint8   `util:"garageID"`
	UserID   uint8   `util:"userID"`
	Model    string  `util:"model"`
	Year     uint8   `util:"year"`
	Price    float64 `util:"price"`
}
