package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type EnvVars struct {
	DB_CONNECTION_STRING string
	ENVIRONMENT          string
}

func LoadEnv() (config EnvVars, err error) {
	err = godotenv.Load()

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	environment := os.Getenv("ENVIRONMENT")

	if err != nil {
		err = errors.New("cannot parse redis DB number")
	}

	config = EnvVars{
		DB_CONNECTION_STRING: connectionString,
		ENVIRONMENT:          environment,
	}

	return config, err
}
