package config

import "gohub/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			"connection": config.Env("DATABASE_CONNECTION"),

			"mysql": map[string]interface{}{
				"host":                 config.Env("DATABASE_HOST", "127.0.0.1"),
				"port":                 config.Env("DATABASE_PORT", "3306"),
				"user":                 config.Env("DATABASE_USER", "root"),
				"password":             config.Env("DATABASE_PASSWORD", "root"),
				"database":             config.Env("DATABASE_NAME", "gohub"),
				"charset":              config.Env("DATABASE_CHARSET", "utf8mb4"),
				"max_idle_connections": config.Env("DATABASE_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": config.Env("DATABASE_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds":     config.Env("DATABASE_MAX_LIFE_SECONDS", 5*60),
			},
		}
	})
}
