package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"score/src/handler"
	"score/src/internal/global"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := global.SetDBToGlobal(); err != nil {
		log.Fatal("Error setting db to global: ", err)
	}

	router := handler.SetUpRouter()
	fmt.Println("Starting server on port", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
