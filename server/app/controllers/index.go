package controllers

import (
	"chat/app/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func AllUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel := models.UserModel()
	users, err := userModel.WhereNe("_id", _id)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)

}

type setAvatarRequestPayload struct {
	Image string `bson:"image"`
}

func SetAvatar(ctx *gin.Context) {
	id := ctx.Param("id")
	var request setAvatarRequestPayload
	// get the request body and bind it to the User object
	if err := ctx.Bind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := map[string]any{
		"avatarImage": request.Image,
	}

	err := models.UserModel().Update(data, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"msg":   "Avatar updated successfully.",
		"image": data["avatarImage"],
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

	messageModel := models.MessageModel()
	log.Println(request)
	senderObjectId, err := primitive.ObjectIDFromHex(request.From)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var data models.MessageWrite
	data.Message = request.Message
	data.Users = []string{
		request.From,
		request.To,
	}
	data.Sender = senderObjectId

	err = messageModel.Create(data)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
	messageModel := models.MessageModel()

	users := []string{
		request.From,
		request.To,
	}

	messages, err := messageModel.WhereEq("users", users)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}
