package entity

type Garage struct {
	GarageID int64  `json:"-" db:"garage_id"`
	UserID   int64  `json:"userId" db:"user_id"`
	Name     string `json:"name" db:"name"`
	Capacity int64  `json:"capacity" db:"capacity"`
}
