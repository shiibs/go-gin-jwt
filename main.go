package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shiibs/go-gin-jwt/database"
	"github.com/shiibs/go-gin-jwt/routes"
)

func init(){
	if err := godotenv.Load(); err !=nil {
		log.Println("Error in loading env file")
	}

	database.ConnectDB()
}

func main(){

	// Close db connection using defer
	psqlDB, err := database.DBConn.DB()

	if err != nil {
		log.Println("Error in getting db connection")
	}

	defer psqlDB.Close()

	port := os.Getenv("port")

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	
	// Add middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Auth-token", "token", "Content-type"},
	}))

	routes.SetupRoutes(router)

	log.Fatal(router.Run(":" + port))
}

