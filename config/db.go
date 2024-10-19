package config

import "github.com/leinodev/core/internal/models/db"

type (
	DBConfig struct {
		Connection Connection `json:"connection"`
		Schemas    []Schema   `json:"schemas"`
	}
	Connection struct {
		Dialect  string              `json:"dialect"`
		Postgres *PostgresConnection `json:"postgres" yaml:"postgres"`
		// TODO: add configuration for connection according to selected dialect
	}
	PostgresConnection struct {
		Host     string `yaml:"host" env:"POSTGRES_HOST"`
		Port     string `yaml:"port" env:"POSTGRES_PORT"`
		User     string `yaml:"user" env:"POSTGRES_USER"`
		DBName   string `yaml:"dbName" env:"POSTGRES_DBNAME"`
		SSLMode  bool   `yaml:"sslMode" env:"SSL_MODE"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	}
	Schema struct {
		Name    string      `json:"name"`
		Tables  []db.Entity `json:"tables"`
		Indexes []db.Index  `json:"indexes"`
	}
)

func (cfg DBConfig) GetConnectionString() string {
	switch cfg.Connection.Dialect {
	case "postgres":
		// TODO: build connection string for Postgres
		return ""
	default:
		return ""
	}
}
