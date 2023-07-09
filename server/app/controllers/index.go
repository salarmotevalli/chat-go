package controllers

import (
	"chat/app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Index(ctx *gin.Context) {
	var userModel models.Query = models.UserModel()
	users, err := userModel.All()

	if err != nil {
		log.Println(err.Error())
	}

	ctx.JSON(http.StatusAccepted, users)
}
