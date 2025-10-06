// sentiric-vertical-hospitality-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort string
	HttpPort string
	CertPath string
	KeyPath  string
	CaPath   string
	LogLevel string
	Env      string

	// Hospitality servisi bağımlılıkları (Placeholder)
	BookingAPIKey  string
	BookingAdapter string
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Dikey Servisler, 201XX bloğu atandı)
	return &Config{
		GRPCPort: GetEnv("VERTICAL_HOSPITALITY_SERVICE_GRPC_PORT", "20111"),
		HttpPort: GetEnv("VERTICAL_HOSPITALITY_SERVICE_HTTP_PORT", "20110"),

		CertPath: GetEnvOrFail("VERTICAL_HOSPITALITY_SERVICE_CERT_PATH"),
		KeyPath:  GetEnvOrFail("VERTICAL_HOSPITALITY_SERVICE_KEY_PATH"),
		CaPath:   GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel: GetEnv("LOG_LEVEL", "info"),
		Env:      GetEnv("ENV", "production"),

		BookingAdapter: GetEnv("BOOKING_ADAPTER", "custom_api"),
		BookingAPIKey:  GetEnv("BOOKING_API_KEY", ""),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}
