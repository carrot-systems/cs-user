package config

import env "github.com/carrot-systems/csl-env"

type GinConfig struct {
	Host string
	Port int
	Mode string
	Tls  bool
}

func LoadGinConfiguration() GinConfig {
	return GinConfig{
		Host: env.RequireEnvString("GIN_LISTEN_URL"),
		Port: env.RequireEnvInt("GIN_PORT"),
		Mode: env.RequireEnvString("GIN_MODE"),
		Tls:  env.RequireEnvBool("GIN_TLS"),
	}
}

