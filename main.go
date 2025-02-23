package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/roh4nyh/qube_challenge_2016/routes"
	"github.com/roh4nyh/qube_challenge_2016/service"
	"github.com/roh4nyh/qube_challenge_2016/utils"
)

func init() {
	if os.Getenv("ENV") == "production" {
		log.Printf("using .env production variables")
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		log.Printf("using .env development variables")
	}

	locations, err := utils.LoadCities()
	if err != nil {
		log.Fatalf("error loading cities: %v", err)
	}

	log.Println("cities loaded successfully")
	log.Printf("total cities: %d", len(locations))

	service.InitDistributorCollection()
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app := gin.Default()
	app.Use(gin.Logger())

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "server is up and running..."})
	})

	// distributor routes
	routes.DistributorRoutes(app)

	log.Printf("server is running on port %s", PORT)
	app.Run(fmt.Sprintf(":%s", PORT))
}
