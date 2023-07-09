package routes

import (
	"chat/app/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	engine := gin.New()

	setRoutes(engine)
	setAuthRoutes(engine)

	return engine
}

func setRoutes(eng *gin.Engine) {
	eng.GET("/", controllers.Index)
	eng.GET("/allusers/:id", controllers.AllUsers)
	eng.POST("setavatart/:id", controllers.SetAvatar)
	eng.POST("addmsg", controllers.AddMessage)
	eng.POST("getmsg", controllers.GetMessage)
}
