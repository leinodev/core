package config

type DbConfig struct {
	ConnectionDSN string `env:"DSN" env-required:""`
}
