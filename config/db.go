package config

import "github.com/leinodev/core/internal/models/db"

type (
	DBConfig struct {
		Connection Connection `json:"connection"`
		Schemas    []Schema   `json:"schemas"`
	}
	Connection struct {
		Dialect string `json:"dialect"`
		// TODO: add configuration for connection according to selected dialect
	}
	Schema struct {
		Tables  []db.Entity `json:"tables"`
		Indexes []db.Index  `json:"indexes"`
	}
)
