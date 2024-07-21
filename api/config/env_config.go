package config

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading enviornment variables:", err)
	}
}
