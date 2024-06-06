package main

import (
	"log"
	"os"

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

	routes.SetupRoutes(router)

	log.Fatal(router.Run(":" + port))
}

