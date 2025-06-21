package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerApi ServerApi
}

type ServerApi struct {
	Port string 
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
	}
}
