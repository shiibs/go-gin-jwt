package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shiibs/go-gin-jwt/controller"
)


func SetupRoutes(r *gin.Engine){

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	


}