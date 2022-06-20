package dto

type CreateCarDto struct {
	GarageID uint8   `json:"garageID"`
	UserID   uint8   `json:"userID"`
	Model    string  `json:"model"`
	Year     uint8   `json:"year"`
	Price    float64 `json:"price"`
}

type UpdateCarDto struct {
	CarID    uint8   `json:"carID"`
	GarageID uint8   `json:"garageID"`
	UserID   uint8   `json:"userID"`
	Model    string  `json:"model"`
	Year     uint8   `json:"year"`
	Price    float64 `json:"price"`
}
