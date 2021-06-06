package getEnv

import (
	"log"
	"os"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
