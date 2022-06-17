package entity

type Car struct {
	CarID    uint8   `util:"-"`
	UserID   uint8   `util:"-"`
	GarageID uint8   `util:"garageID"`
	Model    string  `util:"model"`
	Year     uint8   `util:"year"`
	Price    float64 `util:"price"`
}
