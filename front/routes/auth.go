package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupAuth(eng *gin.Engine) {
	eng.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, "logged in")
	})

	eng.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, "registered")
	})
}
