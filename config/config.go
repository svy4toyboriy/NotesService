package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PgURL      string
	PgProto    string
	PgAddr     string
	PgDb       string
	PgUser     string
	PgPassword string
}

// Get читает из файла переменных среды
func Get() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		PgURL:      os.Getenv("PG_URL"),
		PgProto:    os.Getenv("PG_PROTO"),
		PgAddr:     os.Getenv("PG_ADDR"),
		PgDb:       os.Getenv("PG_DB"),
		PgUser:     os.Getenv("PG_USER"),
		PgPassword: os.Getenv("PG_PASSWORD"),
	}
}
