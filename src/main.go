package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"score/src/handler/routes"
)

func main() {
	//TODO: Add GOPATH
	if err := godotenv.Load("/Users/yogendravutukuri/GolandProjects/go/src/score/src/.env"); err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(-1)
	}
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/tournaments", routes.GetTournaments)
	fmt.Println("Starting server on port", os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
