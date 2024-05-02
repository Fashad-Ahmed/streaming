package env

import "github.com/joho/godotenv"

func EnvConfigurationHandler() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
