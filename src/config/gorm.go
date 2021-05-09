package config

import env "github.com/carrot-systems/csl-env"

type GormConfig struct {
	Engine   string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func LoadGormConfiguration() GormConfig {
	return GormConfig{
		Engine:   env.RequireEnvString("DB_ENGINE"),
		Host:     env.RequireEnvString("DB_HOST"),
		Port:     env.RequireEnvInt("DB_PORT"),
		User:     env.RequireEnvString("DB_USER"),
		Password: env.RequireEnvString("DB_PASSWORD"),
		DbName:   env.RequireEnvString("DB_NAME"),
	}
}
