package routes

import (
	"chat/app/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	engine := gin.New()
	api := engine.Group("/api")
	setMessageRoutes(api)
	setAuthRoutes(api)

	return engine
}

func setMessageRoutes(api *gin.RouterGroup) {
	// Set prefix
	msg := api.Group("/messages")

	msg.GET("/", controllers.Index)
	msg.GET("/allusers/:id", controllers.AllUsers)
	msg.POST("/setavatart/:id", controllers.SetAvatar)
	msg.POST("/addmsg", controllers.AddMessage)
	msg.POST("/getmsg", controllers.GetMessage)
}
