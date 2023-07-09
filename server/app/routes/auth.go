package routes

import (
	"chat/app/controllers"
	"github.com/gin-gonic/gin"
)

func setAuthRoutes(eng *gin.Engine) {
	eng.POST("/login", controllers.Login)
	eng.POST("/register", controllers.Register)
	eng.POST("/logout", controllers.LogOut)
}
