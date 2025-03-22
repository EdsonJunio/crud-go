package controller

import (
	"fmt"
	"github.com/EdsonJunio/crud-go/src/configuration/rest_err"
	"github.com/EdsonJunio/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserResquest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadResquestErro(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
}
