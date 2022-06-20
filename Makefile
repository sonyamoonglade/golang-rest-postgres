

migr:
	migrate -path migrations -database "postgresql://postgres:admin@localhost:5432/golang?sslmode=disable" -verbose up

