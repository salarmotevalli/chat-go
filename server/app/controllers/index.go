package controllers

import (
	"chat/app/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type getMessagePayload struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func GetMessage(ctx *gin.Context) {
	var request getMessagePayload

	// get the request body and bind it to the User object
	if err := ctx.Bind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageModel := models.MessageModel()
	from, err := primitive.ObjectIDFromHex(request.From)
	to, err := primitive.ObjectIDFromHex(request.To)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users := []primitive.ObjectID{
		from,
		to,
	}

	messages, err := messageModel.WhereEq("users", users)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)

}
