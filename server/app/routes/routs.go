package routes

import (
	"chat/app/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	engine := gin.New()
	engine.Use(corsMiddleware())
	api := engine.Group("/api")
	setMessageRoutes(api)
	setAuthRoutes(api)

	return engine
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func setMessageRoutes(api *gin.RouterGroup) {
	// Set prefix
	msg := api.Group("/messages")

	msg.POST("/setavatart/:id", controllers.SetAvatar)
	msg.POST("/addmsg", controllers.AddMessage)
	msg.POST("/getmsg", controllers.GetMessage)
}
