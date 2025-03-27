package cfg

import (
	"log"
	"os"
)

func getEnv(key, fallback string) string {
	val, found := os.LookupEnv(key)
	if !found {
		log.Printf("%s environment variable not found! using fallback.\n", key)
		return fallback
	}
	return val
}

func mustGetEnv(key string) string {
	val, found := os.LookupEnv(key)
	if !found {
		log.Fatalf("%s environment variable not found, set variable first!\n", key)
	}
	return val
}
