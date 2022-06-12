package entities

type User struct {
	UserID   uint8   `json:"-"`
	GarageID uint8   `json:"garageId"`
	Name     string  `json:"name"`
	Salary   float64 `json:"salary"`
}
