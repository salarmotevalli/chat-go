package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	engine := gin.New()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, "Hello, world!")
	})

	setupAuth(engine)

	return engine
}
