package entities

type Car struct {
	CarID    uint8   `json:"-"`
	UserID   uint8   `json:"-"`
	GarageID uint8   `json:"garageId"`
	Model    string  `json:"model"`
	Year     uint8   `json:"year"`
	Price    float64 `json:"price"`
}
