package config

type RedisConfig struct {
	Hostname string `env:"HOSTNAME" env-required:""`
	Password string `env:"PASSWORD" env-required:""`
	DbNum    int    `env:"DBNUM" env-required:""`
}
