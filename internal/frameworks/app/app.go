package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"storage/internal/frameworks/database"
	"storage/internal/frameworks/di"
)

func Setup() *gin.Engine {
	loadEnvFile()

	routes := gin.Default()
	db := database.NewSession()

	dependencyInjection := di.NewDependencyInjection(routes, db)
	dependencyInjection.SetupDependencies()

	return routes
}

func loadEnvFile() {
	err := godotenv.Load("")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
