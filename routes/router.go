package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shiibs/go-gin-jwt/controller"
	"github.com/shiibs/go-gin-jwt/middleware"
)


func SetupRoutes(r *gin.Engine){

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)

	private := r.Group("/private")

	private.Use(middleware.Authenticate)

	private.GET("/refreshToken", controller.RefreshToken)
	


}