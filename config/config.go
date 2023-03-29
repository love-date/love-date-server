package config

import (
	"os"
	"strconv"
)

type HttpServerConfig struct {
	Host string
	Port int
}

type SqlDataBase struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type Config struct {
	Server HttpServerConfig
	SqlDB  SqlDataBase
}

func New() *Config {
	return &Config{
		Server: HttpServerConfig{
			Host: getEnv("SERVER_HOST", ""),
			Port: getEnvAsInt("SERVER_PORT", 1988),
		},
		SqlDB: SqlDataBase{
			Username: getEnv("DB_USERNAME", "masood"),
			Password: getEnv("DB_PASSWORD", ""),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			Name:     getEnv("DB_NAME", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
