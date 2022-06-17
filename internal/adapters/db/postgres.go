package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/config"
)

func GetDbInstance(cfg *config.DbConfig) (*sqlx.DB, error) {

	db, err := sqlx.Open(cfg.Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.DatabaseName, cfg.Password))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
