package main

import (
	"log"
	"os"

	"login/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run(":" + os.Getenv("PORT"))
}
