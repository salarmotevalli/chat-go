package routes

import (
	"chat/app/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	engine := gin.New()

	engine.GET("/", controllers.Index)

	setupAuth(engine)

	return engine
}
