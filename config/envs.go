package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Envinroments struct {
	Port           string
	SecretKey      string
	DbURL          string
	ExpirationTime int
	RedisAddr      string
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
	} else {
		Envs.Port = port
	}

	expirationTime := os.Getenv("EXPIRATION_TIME")
	if expirationTime == "" {
		Envs.ExpirationTime = 60
	} else {
		Envs.ExpirationTime, _ = strconv.Atoi(expirationTime)
	}

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

	redisAddr, ok := os.LookupEnv("REDIS_ADDR")
	if !ok || redisAddr == "" {
		return errors.New("REDIS_ADDR not set")
	}
	Envs.RedisAddr = redisAddr

	return nil
}
