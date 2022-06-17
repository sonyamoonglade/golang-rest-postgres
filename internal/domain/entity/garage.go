package entity

type Garage struct {
	GarageID uint8  `util:"-"`
	Name     string `util:"name"`
	Capacity uint8  `util:"capacity"`
}
