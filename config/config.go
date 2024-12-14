package config

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	} else {
		log.Printf("Environment variable %s loaded successfully", key)
	}

	return value
}
