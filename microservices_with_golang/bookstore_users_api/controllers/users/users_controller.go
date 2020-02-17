package users

import (
	"../../domain/users"
	"../../services"
	"../../utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"

)

func CreateUser(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: return bad request to the called
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me!")
}