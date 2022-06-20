package dto

type CreateGarageDto struct {
	UserID   int64  `json:"-"`
	Name     string `json:"name"`
	Capacity int64  `json:"capacity"`
}
