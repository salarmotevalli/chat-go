package routes

import (
	"chat/app/controllers"
	"github.com/gin-gonic/gin"
)

func setAuthRoutes(api *gin.RouterGroup) {
	// Set prefix
	auth := api.Group("auth")

	auth.POST("/setavatar/:id", controllers.SetAvatar)
	auth.GET("/allusers/:id", controllers.AllUsers)
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)
	auth.GET("/logout/:id", controllers.LogOut)
}
