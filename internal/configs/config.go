package configs

import (
	"test-1/pkg/env"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type ServerConfig struct {
	Port    int
	Address string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     env.GetEnv("DB_HOST", "localhost"),
			Port:     env.GetEnvAsInt("DB_PORT", 5432),
			Username: env.GetEnv("DB_USER", "user"),
			Password: env.GetEnv("DB_PASSWORD", "password"),
			DBName:   env.GetEnv("DB_NAME", "appdb"),
		},
		Server: ServerConfig{
			Port:    env.GetEnvAsInt("SERVER_PORT", 8080),
			Address: env.GetEnv("SERVER_ADDRESS", "0.0.0.0"),
		},
	}, nil
}
