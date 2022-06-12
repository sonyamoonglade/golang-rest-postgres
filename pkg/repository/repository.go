package repository

type Repository struct {
	CarRepository
}

type CarRepository interface {
}

func NewRepository() *Repository {
	//
	//DbConfig{
	//	Dialect: "postgres",
	//	DBName: ""
	//}

	return &Repository{}
}
