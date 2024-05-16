package main

import (
	"fmt"
	"log"
	"os"

	"assalam-learn/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error read env file with err: %s", err)
	}

	e := routes.Routes()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
