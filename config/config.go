package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PerIP              int
	BlockDurationIP    time.Duration
	PerToken           int
	BlockDurationToken time.Duration
	RedisHost          string
	RedisPassword      string
	RedisDB            int
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		PerIP:              getEnvAsInt("RATE_LIMIT_PER_IP", 5),
		BlockDurationIP:    getEnvAsDuration("RATE_LIMIT_BLOCK_DURATION_IP", 5*time.Minute),
		PerToken:           getEnvAsInt("RATE_LIMIT_PER_TOKEN", 10),
		BlockDurationToken: getEnvAsDuration("RATE_LIMIT_BLOCK_DURATION_TOKEN", 5*time.Minute),
		RedisHost:          getEnvOrDefault("REDIS_HOST", "localhost:6379"),
		RedisPassword:      getEnvOrDefault("REDIS_PASSWORD", ""),
		RedisDB:            getEnvAsInt("REDIS_DB", 0),
	}
}
func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
func getEnvAsInt(key string, defaultVal int) int {
	if valStr := os.Getenv(key); valStr != "" {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}

func getEnvAsDuration(key string, defaultVal time.Duration) time.Duration {
	if valStr := os.Getenv(key); valStr != "" {
		if val, err := time.ParseDuration(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}
