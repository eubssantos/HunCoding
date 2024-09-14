package controller

import (
	"HunCoding/src/main/controller/model/request"
	"HunCoding/src/main/configuration/logger"
	"HunCoding/src/main/configuration/validation"
	"HunCoding/src/main/model"
	"HunCoding/src/main/model/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", 
			zap.String("journey", "createUser"),
	)		

	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, 
				zap.String("journey", "createUser"))	
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully", 
			zap.String("journey", "createUser"),
	)

	c.String(http.StatusOK, "")
}
