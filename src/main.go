package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"score/src/handler"
	"score/src/internal/global"
)

func main() {
	//TODO: Add GOPATH
	if err := godotenv.Load("/Users/yogendravutukuri/GolandProjects/go/src/score/src/.env"); err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(-1)
	}

	//Acquiring DB connection
	if err := global.SetDBToGlobal(); err != nil {
		fmt.Println("Error setting db to global", err)
		os.Exit(-1)
	}

	//Using GIN package initiating routes
	router := handler.SetUpRouter()
	fmt.Println("Starting server on port", os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
