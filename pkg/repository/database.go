package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Dialect  string
}

func GetDbInstance(cfg DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Dialect, fmt.Sprintf("port=%s user=%s host=%s password=%s sslmode=disable dbname=%s\n", cfg.Port, cfg.Username, cfg.Host, cfg.Password, cfg.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
