package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func httpRequest() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf(string(body))
}

func main() {
	API_TOKEN := goDotEnvVariable("DISCORD_API_TOKEN")

	fmt.Printf(API_TOKEN)
}
