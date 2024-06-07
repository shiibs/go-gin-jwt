package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shiibs/go-gin-jwt/database"
	"github.com/shiibs/go-gin-jwt/helper"
	"github.com/shiibs/go-gin-jwt/model"
	"golang.org/x/crypto/bcrypt"
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

	// Check user credential 
	var formData formData

	if err := c.ShouldBind(&formData); err != nil {
		log.Println("Error in json binding")
		returnObject["mgs"] = "Error in form data"
		c.JSON(400, returnObject)
		return
	}

	var user model.User

	database.DBConn.First(&user, "email=?", formData.Email)

	if user.ID == 0 {
		returnObject["msg"] = "User not found"
		
		c.JSON(400, returnObject)
		return 
	}

	// Validate password 
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))


	if err != nil {
		log.Println("Invalid password.")

		returnObject["msg"] = "Invalid password"
		c.JSON(401, returnObject)
		return
	}

	// create jwt token

	token, err := helper.GenerateToken(user)

	if err != nil {
		returnObject["msg"] = "Token creation error."
		c.JSON(401, returnObject)
		return
	}

	returnObject["token"] = token
	returnObject["user"] = user
	returnObject["status"] = "OK"
	returnObject["msg"] = "User authenticated"
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

func RefreshToken(c *gin.Context){
	returnObject := gin.H{
		"status": "OK",
		"msg":    "Refresh Token  route",
	}

	email, exists := c.Get("email")

	if !exists {
		log.Println("Email not found")
		returnObject["msg"] = "Email not found."
		c.JSON(401, returnObject)
		return
	}

	var user model.User 

	database.DBConn.First(&user, "email=?", email)

	if user.ID == 0 {
		log.Println("User not found.")
		returnObject["msg"] = "User not found."

		c.JSON(400, returnObject)
		return
	}

	token, err := helper.GenerateToken(user)

	if err != nil {
		returnObject["msg"] = "Token creation error."
		c.JSON(401, returnObject)
		return
	}

	returnObject["token"] = token
	returnObject["user"] = user

	c.JSON(200, returnObject)
}