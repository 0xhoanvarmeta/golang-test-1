package env

import (
	"os"
	"strconv"
)

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		if defaultValue == "" {
			panic("Environment variable " + key + " not set")
		}

		return defaultValue
	}

	return value
}

func GetEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)

	if value == "" {
		if defaultValue == 0 {
			panic("Environment variable " + key + " not set")
		}

		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}
