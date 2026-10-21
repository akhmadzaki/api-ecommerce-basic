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
	ServerPort        string
	DB                DBConfig
	JWTAccessSecret   string
	JWTRefreshSecret  string
	DefaultAPITimeout string
}

func LoadConfig() (*Config, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("DB_HOST environment variable is required")
	}

	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		return nil, fmt.Errorf("POSTGRES_DB environment variable is required")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("POSTGRES_USER environment variable is required")
	}

	dbPass := os.Getenv("POSTGRES_PASSWORD")
	if dbPass == "" || len(dbPass) < 32 {
		return nil, fmt.Errorf("strong POSTGRES_PASSWORD environment variable is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, fmt.Errorf("DB_PORT environment variable is required")
	}

	jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")
	if jwtAccessSecret == "" || len(jwtAccessSecret) < 32 {
		return nil, fmt.Errorf("strong JWT_ACCESS_SECRET environment variable is required")
	}

	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	if jwtRefreshSecret == "" || len(jwtRefreshSecret) < 32 {
		return nil, fmt.Errorf("strong JWT_REFRESH_SECRET environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default fallback
	}

	defaultApiTimeout := os.Getenv("DEFAULT_API_TIMEOUT")
	if defaultApiTimeout == "" {
		defaultApiTimeout = "30" // Default fallback
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
