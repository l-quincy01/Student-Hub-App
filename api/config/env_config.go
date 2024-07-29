package config

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(filepath.Join("..", "..", ".env"))
	if err != nil {
		log.Fatal("error loading enviornment variables:", err)
	}
}
