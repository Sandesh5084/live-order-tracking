package config

import (
	"os"
	
)

type Config struct {
	Port  string
	DBURL string
}

func LoadConfig() Config {
	return Config{
		Port:  GetEnv("PORT", "8080"),
		DBURL: GetEnv("DATABASE_URL", "postgres://lot_user:lot_pass@localhost:5432/lot_db?sslmode=disable"),
	}
}

func GetEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
