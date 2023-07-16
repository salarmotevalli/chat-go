package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chat/app/services"
)

type loginRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var request loginRequestPayload

	// Bind whit payload
	if err := ctx.Bind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.LoadUserService().UserAuthenticate(request.Username, request.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, map[string]any{
		"status": true,
		"user":   user,
	})

}

type registerRequestPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context) {
	var request registerRequestPayload

	// Bind whit payload
	if err := ctx.Bind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.LoadUserService().CreateUser(request.Username , request.Email, request.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status": true,
		"user":   user,
	})
}

func LogOut(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
