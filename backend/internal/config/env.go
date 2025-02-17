package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	Port    string
	Prefix  string
	Version string

	DbPort string
	DbName string
	DbHost string
	DbUser string
	DbPwd  string
	DbSsl  string
}

func LoadEnv() *ConfigEnv {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env file: %v", err)
	}

	return &ConfigEnv{
		Port:    os.Getenv("PORT"),
		Prefix:  os.Getenv("PREFIX"),
		Version: os.Getenv("VERSION"),

		DbPort: os.Getenv("DB_PORT"),
		DbHost: os.Getenv("DB_HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPwd:  os.Getenv("DB_PASSWORD"),
		DbSsl:  os.Getenv("DB_SSL"),
	}
}
