package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load(envFile string) error {
	return godotenv.Load(envFile)
}

func Get(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) (int, error) {
	value := os.Getenv(key)
	return strconv.Atoi(value)
}
