package config

import (
	"flag"
	"fmt"
)

const (
	dbDSN = "host=%s port=5432 dbname=postgres user=admin password=admin sslmode=disable"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}

func NewPGConfig() (PGConfig, error) {
	var dbHost string
	flag.StringVar(&dbHost, "dbhost", "localhost", "database connection host")
	flag.Parse()

	return &pgConfig{
		dsn: fmt.Sprintf(dbDSN, dbHost),
	}, nil
}
