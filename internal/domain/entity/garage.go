package entity

type Garage struct {
	GarageID uint8  `json:"-"`
	Name     string `json:"name"`
	Capacity uint8  `json:"capacity"`
}
