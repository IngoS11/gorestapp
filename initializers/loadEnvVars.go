package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Did not find .env file")
	}

	database_dsn := os.Getenv("DATABASE_DSN")
	if database_dsn == "" {
		log.Fatal("DATABASE_DSN missing")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		log.Fatal("JWT_SECRET missing")
	}

}
