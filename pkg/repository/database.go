package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Dialect  string
}

func GetDbInstance(cfg DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Dialect, fmt.Sprintf("port=%s user=%s host=%s dbname=%s\n", cfg.Port, cfg.Username, cfg.Host, cfg.DBName))
	if err != nil {
		log.Fatalf("error occured when connecting to databae... %s", err.Error())
	}
	return db, nil
}
