package dto

type CreateCarDto struct {
	Model string  `json:"model"`
	Year  int64   `json:"year"`
	Price float64 `json:"price"`
}

type UpdateCarDto struct {
	CarID int64   `json:"carId"`
	Model string  `json:"model,omitempty"`
	Year  int64   `json:"year,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type BuyCarDto struct {
	UserID   int64 `json:"-"`
	CarID    int64 `json:"carId"`
	GarageID int64 `json:"garageId"`
}
