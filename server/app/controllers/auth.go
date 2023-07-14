package controllers

import (
	"chat/app/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

	user, err := models.UserModel().FindField("username", request.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("username or email is incorrect")})
		return
	}

	if comparePasswords(request.Password, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("username or email is incorrect")})
		return
	}

	user.Password = ""

	ctx.JSON(http.StatusCreated, user)

}

func comparePasswords(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func Register(ctx *gin.Context) {

}

func LogOut(ctx *gin.Context) {

}
