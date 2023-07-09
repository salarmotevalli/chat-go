package controllers

import (
	"chat/app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Index(ctx *gin.Context) {
	var messageModel models.Query = models.MessageModel()
	messages, err := messageModel.All()

	if err != nil {
		log.Println(err.Error())
	}

	ctx.JSON(http.StatusAccepted, messages)
}

func AllUsers(ctx *gin.Context) {

}

func SetAvatar(ctx *gin.Context) {

}

func AddMessage(ctx *gin.Context) {

}

func GetMessage(ctx *gin.Context) {

}
