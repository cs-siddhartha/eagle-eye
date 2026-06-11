package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	defaultAddress           = ":8080"
	defaultMaxRequestBytes   = int64(1 << 20)
	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 15 * time.Second
	defaultWriteTimeout      = 15 * time.Second
	defaultIdleTimeout       = 60 * time.Second
	defaultShutdownTimeout   = 10 * time.Second
)

type Config struct {
	Address           string
	MaxRequestBytes   int64
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ShutdownTimeout   time.Duration
}

// Load reads supported environment variables and rejects unsafe values before
// the HTTP server starts accepting traffic.
func Load() (Config, error) {
	config := Config{
		Address:           envOrDefault("API_ADDRESS", defaultAddress),
		MaxRequestBytes:   defaultMaxRequestBytes,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		IdleTimeout:       defaultIdleTimeout,
		ShutdownTimeout:   defaultShutdownTimeout,
	}

	if value, ok := os.LookupEnv("API_MAX_REQUEST_BYTES"); ok {
		maxRequestBytes, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return Config{}, fmt.Errorf("API_MAX_REQUEST_BYTES must be a valid integer: %w", err)
		}
		if maxRequestBytes <= 0 {
			return Config{}, fmt.Errorf("API_MAX_REQUEST_BYTES must be greater than zero")
		}

		config.MaxRequestBytes = maxRequestBytes
	}

	durationSettings := []struct {
		key    string
		target *time.Duration
	}{
		{key: "API_READ_HEADER_TIMEOUT", target: &config.ReadHeaderTimeout},
		{key: "API_READ_TIMEOUT", target: &config.ReadTimeout},
		{key: "API_WRITE_TIMEOUT", target: &config.WriteTimeout},
		{key: "API_IDLE_TIMEOUT", target: &config.IdleTimeout},
		{key: "API_SHUTDOWN_TIMEOUT", target: &config.ShutdownTimeout},
	}

	for _, setting := range durationSettings {
		value, ok := os.LookupEnv(setting.key)
		if !ok {
			continue
		}

		duration, err := time.ParseDuration(value)
		if err != nil {
			return Config{}, fmt.Errorf("%s must be a valid duration: %w", setting.key, err)
		}
		if duration <= 0 {
			return Config{}, fmt.Errorf("%s must be greater than zero", setting.key)
		}

		*setting.target = duration
	}

	return config, nil
}

// envOrDefault keeps optional string configuration concise and predictable.
func envOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	return fallback
}
