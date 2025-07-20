package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerApi ServerApi
	Postgres Postgres
}

type ServerApi struct {
	Port string
}

type Postgres struct {
	Host string
	Port string
	Username string
	Password string
	Database string
}

func InitConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		ServerApi: ServerApi{
			Port: os.Getenv("API_SERVER_PORT"),
		},
		Postgres: Postgres{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Username: os.Getenv("POSTGRES_USERNAME"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DATABASE"),
		},
	}
}
