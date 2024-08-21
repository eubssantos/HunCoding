package controller

import (
	"HunCoding/src/main/command/user/request"
	"HunCoding/src/main/configuration/validation"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUse controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error %s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
}
