package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"chat/app/models"
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("username or email is incorrect").Error()})
		return
	}

	if comparePasswords(request.Password, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("username or email is incorrect").Error()})
		return
	}

	user.Password = ""

	ctx.JSON(http.StatusCreated, user)

}

func comparePasswords(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
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
	userModel := models.UserModel()
	user, _ := userModel.FindField("email", request.Email)

	if user != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("user already exist").Error()})
		return
	}

	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	data := models.UserWrite{
		Email:    request.Email,
		Password: string(hashBytes),
		Username: request.Username,
	}

	createdUser, err := userModel.Create(data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// createdUser := userModel.FindId(createdUser)




	ctx.JSON(http.StatusCreated, map[string]any{
		"status": true,
		"user": createdUser,
	})
}

func LogOut(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
