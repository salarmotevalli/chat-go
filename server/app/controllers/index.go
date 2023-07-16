package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"chat/app/services"
)

func AllUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	
	// Get all users except logged in user
	users, err := services.LoadUserService().GetOtherUsers(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type setAvatarRequestPayload struct {
	Image string `json:"image"`
}

func SetAvatar(ctx *gin.Context) {
	id := ctx.Param("id")
	var request setAvatarRequestPayload

	// Get the request body and bind it to the User object
	if err := ctx.Bind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.LoadUserService().UpdateUserAvatar(request.Image, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"isSet": true,
		"image": request.Image,
	})
}

type addMessageRequestPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func AddMessage(ctx *gin.Context) {
	var request addMessageRequestPayload

	// get the request body and bind it to the User object
	if err := ctx.Bind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.LoadMessageService().CreateMessage(request.From, request.To, request.Message)
	if err!= nil {
		
	}

	ctx.JSON(http.StatusCreated, map[string]string{
		"msg": "Message added successfully.",
	})
}

type getMessageRequestPayload struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func GetMessage(ctx *gin.Context) {
	var request getMessageRequestPayload

	// get the request body and bind it to the User object
	if err := ctx.Bind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	messages, err := services.LoadMessageService().GetMessages(request.From, request.To)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}
