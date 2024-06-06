package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shiibs/go-gin-jwt/database"
)

func init(){
	if err := godotenv.Load(); err !=nil {
		log.Println("Error in loading env file")
	}

	database.ConnectDB()
}

func main(){

	port := os.Getenv("port")

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "HelloWorld",
		})
	})

	log.Fatal(router.Run(":" + port))
}

