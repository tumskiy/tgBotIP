package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("not find .env file")
	}
}

func ParseEnv(key string) string {
	return os.Getenv(key)
}
