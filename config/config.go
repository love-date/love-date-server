package config

import (
	"os"
	"strconv"
)

type httpServerConfig struct {
	Host string
	Port int
}

type sqlDataBase struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type jwt struct {
	Key string
}

type Config struct {
	Server httpServerConfig
	SqlDB  sqlDataBase
	Jwt    jwt
}

func New() *Config {
	return &Config{
		Server: httpServerConfig{
			Host: getEnv("SERVER_HOST", ""),
			Port: getEnvAsInt("SERVER_PORT", 1988),
		},
		SqlDB: sqlDataBase{
			Username: getEnv("DB_USERNAME", "masood"),
			Password: getEnv("DB_PASSWORD", ""),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			Name:     getEnv("DB_NAME", ""),
		},
		Jwt: jwt{
			Key: getEnv("JWT_KEY", "test123jwt"),
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
