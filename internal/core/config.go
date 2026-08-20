package core

import (
	"os"
	"strconv"
	"strings"
)

// Config is loaded once at startup and stays immutable during a run.
type Config struct {
	Port int
	Host string
	Verbose bool
	TimeoutS int
	Retries int
	LogLevel string
	DataDir string
}

// LoadConfig reads configuration from environment variables.
func LoadConfig() Config {
	return Config{
 Port: envInt("PORT", 8080),
 Host: envStr("HOST", "127.0.0.1"),
 Verbose: envBool("VERBOSE"),
 TimeoutS: envInt("TIMEOUT_S", 30),
 Retries: envInt("RETRIES", 3),
 LogLevel: strings.ToLower(envStr("LOG_LEVEL", "info")),
 DataDir: envStr("DATA_DIR", "./data"),
	}
}

func envInt(key string, fallback int) int {
	raw := os.Getenv(key)
	if raw == "" {
 return fallback
	}
	if v, err := strconv.Atoi(raw); err == nil {
 return v
	}
	return fallback
}

func envBool(key string) bool {
	return os.Getenv(key) == "1" || os.Getenv(key) == "true"
}

func envStr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
 return v
	}
	return fallback
}