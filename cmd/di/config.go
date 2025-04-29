package di

import "os"

type Env struct {
	// HTTP Service
	HttpPort string
}

func GetEnv() Env {
	return Env{
		HttpPort: getEnvKey("HTTP_PORT", "4000"),
	}
}

func getEnvKey(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
