package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

type Config struct {
	ServerPort       string
	DB               DBConfig
	JWTAccessSecret  string
	JWTRefreshSecret string
}

func LoadConfig() (*Config, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("DB_HOST environment variable is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB_NAME environment variable is required")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("DB_USER environment variable is required")
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		return nil, fmt.Errorf("DB_PASSWORD environment variable is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, fmt.Errorf("DB_PORT environment variable is required")
	}

	jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")
	if jwtAccessSecret == "" {
		return nil, fmt.Errorf("JWT_ACCESS_SECRET environment variable is required")
	}

	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	if jwtRefreshSecret == "" {
		return nil, fmt.Errorf("JWT_REFRESH_SECRET environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default fallback
	}

	return &Config{
		ServerPort: port,
		DB: DBConfig{
			Host:     dbHost,
			Name:     dbName,
			User:     dbUser,
			Password: dbPass,
			Port:     dbPort,
		},
		JWTAccessSecret:  jwtAccessSecret,
		JWTRefreshSecret: jwtRefreshSecret,
	}, nil
}
