package envconfig

import "os"

type EnvConfig struct {
	PgConfig PostgresConfig
}

type PostgresConfig struct {
	User    string
	Pwd     string
	DB      string
	Host    string
	Port    string
	Sslmode string
}

func GetEnv() EnvConfig {
	return EnvConfig{
		PgConfig: PostgresConfig{
			User:    os.Getenv("POSTGRES_USER"),
			Pwd:     os.Getenv("POSTGRES_PWD"),
			DB:      os.Getenv("POSTGRES_DB"),
			Host:    os.Getenv("POSTGRES_HOST"),
			Port:    os.Getenv("POSTGRES_PORT"),
			Sslmode: os.Getenv("POSTGRES_SSLMODE"),
		},
	}
}
