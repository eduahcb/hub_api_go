package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Envinroments struct {
	Port      string
	SecretKey string
	DbURL     string
}

var Envs Envinroments

func initEnvs() error {
	mode := os.Getenv("HUB_API")

	if mode == "" {
		mode = "development"
	}

	if mode == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

  port := os.Getenv("PORT")
  if port == "" {
    Envs.Port = "8080"
  }
  Envs.Port = port

	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok || secretKey == "" {
		return errors.New("SECRET_KEY not set")
	}
	Envs.SecretKey = secretKey


	dbURL, ok := os.LookupEnv("DB_URL")
	if !ok || dbURL == "" {
		return errors.New("DB_URL not set")
	}
	Envs.DbURL = dbURL

	return nil
}
