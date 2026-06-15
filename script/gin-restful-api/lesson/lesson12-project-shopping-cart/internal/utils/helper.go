package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"trungem.com/shopping-cart/pkg/logger"
)

func GetEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {
	val := os.Getenv(key)

	if val == "" {
		return defaultValue
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}

	return intVal
}

func NewLoggerWithPath(fileName, level string) *zerolog.Logger {
	cwd, err := os.Getwd()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Unable to get working dir")
	}

	path := filepath.Join(cwd, "internal/logs", fileName)

	config := logger.LoggerConfig{
		Level:      level,
		Filename:   path,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
		LocalTime:  true,
		IsDev:      GetEnv("APP_ENV", "development"),
	}

	return logger.NewLogger(config)
}

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func MustGetWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Unable to get working dir")
	}

	return dir
}
