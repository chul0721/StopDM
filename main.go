package main

import (
	"/utils/getEnv"
	"fmt"
)

func main() {
	API_TOKEN := getEnv.GoDotEnvVariable("DISCORD_API_TOKEN")

	fmt.Printf(API_TOKEN)
}
