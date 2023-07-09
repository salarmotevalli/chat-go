package controllers

import (
	"chat/app/models"
	"github.com/gin-gonic/gin"
	"log"

	"net/http"
)

func Index(ctx *gin.Context) {
	userModel := models.UserModel()
	users, err := userModel.All()

	if err != nil {
		log.Println(err)
		//return
	}

	ctx.JSON(http.StatusAccepted, users)
}
