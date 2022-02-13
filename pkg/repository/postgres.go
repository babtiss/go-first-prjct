package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

//
// for run:
// docker run --name=go-db -e POSTGRES_PASSWORD='babtis' -p 5436:5432 -d --rm postgres
// migrate -path ./schema -database 'postgres://postgres:babtis@localhost:5436/postgres?sslmode=disable' up
// for check:
// docker exec -it <id> /bin/bash
//
