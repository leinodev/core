package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type LaunchConfig struct {
	AppInstance string `env:"APP_INSTANCE" env-default:""`

	DbConfig    DbConfig    `env-prefix:"DB_"`
	RedisConfig RedisConfig `env-prefix:"REDIS_"`
}

func LoadLaunchConfig() (*LaunchConfig, error) {
	cfg := &LaunchConfig{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
