package config

import "github.com/ilyakaznacheev/cleanenv"

type CoreConfig struct {
	DB DBConfig `json:"db"`
}

func NewConfig(path string) (*CoreConfig, error) {
	config := new(CoreConfig)
	if err := config.readConfig(path); err != nil {
		return nil, err
	}
	return config, nil
}

func (config *CoreConfig) readConfig(path string) error {
	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return err
	}
	return nil
}
