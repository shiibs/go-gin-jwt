package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shiibs/go-gin-jwt/database"
	"github.com/shiibs/go-gin-jwt/helper"
	"github.com/shiibs/go-gin-jwt/model"
)

type formData struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// Login handler
func Login(c *gin.Context){
	returnObject := gin.H{
		"status": "OK",
		"msg": "Login route",
	}
	c.JSON(200, returnObject)

}

// Register handler
func Register(c *gin.Context){
	returnObject := gin.H{
		"status": "OK",
		"msg": "Register route",
	}

	// collect formData
	var formData formData

	if err := c.ShouldBind(&formData); err != nil {
		log.Println("Error in json binding")
		returnObject["mgs"] = "Error in form data"
		c.JSON(400, returnObject)
		return
	}

	// add form data to user model
	var user model.User

	user.Email = formData.Email
	user.Password = helper.HashPassword(formData.Password)

	// Add user to database
	result := database.DBConn.Create(&user)
	
	if result.Error != nil {
		log.Println(result.Error)
		returnObject["msg"] = "User already exists."
		c.JSON(400, returnObject)
		return
	}

	returnObject["msg"] = "User added successfully."
	c.JSON(201, returnObject)
}

func Logout(){}

func RefreshToken(){}