package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Variables struct {
	Host string
	Port string
	DBuser string
	Password string
	DBname string
}

func ProjectEnvs() *Variables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	return &Variables{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DBuser: os.Getenv("DB_USER"),
		Password: os.Getenv("PASSWORD"),
		DBname: os.Getenv("DBNAME"),
	}
}
