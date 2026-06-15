package config

import (
	"fmt"

	"trungem.com/hoc-gin/internal/utils"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	BD DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		BD: DatabaseConfig{
			Host:     utils.GetEnv("DB_HOST", "localhost"),
			Port:     utils.GetEnv("DB_PORT", "5432"),
			User:     utils.GetEnv("DB_USER", "postgres"),
			Password: utils.GetEnv("DB_PASSWORD", "postgres"),
			DBName:   utils.GetEnv("DB_NAME", "master-golang"),
			SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
		},
	}
}

func (c *Config) DNS() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.BD.Host, c.BD.Port, c.BD.User, c.BD.Password, c.BD.DBName, c.BD.SSLMode)
}
