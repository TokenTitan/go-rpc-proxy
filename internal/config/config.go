package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
	Services        map[string]string
}

func Load() *Config {
	return &Config{
		Port:            getEnv("PROXY_PORT", "8080"),
		ShutdownTimeout: getDuration("SHUTDOWN_TIMEOUT_SEC", 10),
		Services: map[string]string{
			"user-service":  getEnv("USER_SERVICE_ADDR", "localhost:9001"),
			"order-service": getEnv("ORDER_SERVICE_ADDR", "localhost:9002"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getDuration(key string, fallbackSec int) time.Duration {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return time.Duration(n) * time.Second
		}
	}
	return time.Duration(fallbackSec) * time.Second
}
