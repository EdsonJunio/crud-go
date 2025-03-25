package controller

import (
	"fmt"
	"github.com/EdsonJunio/crud-go/src/configuration/validation"
	"github.com/EdsonJunio/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser Controller")
	var userRequest request.UserResquest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Println("Error trying to marshal object, error%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
