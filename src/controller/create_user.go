package controller

import (
	"github.com/EdsonJunio/crud-go/src/configuration/validation"
	"github.com/EdsonJunio/crud-go/src/controller/model/request"
	"github.com/EdsonJunio/crud-go/src/controller/model/response"
	"github.com/HunCoding/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserResquest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validat user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User create successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusCreated, response)
}
