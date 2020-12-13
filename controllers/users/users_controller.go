package users

import (
	"github.com/PMerdala/users-api/domain/users"
	users_service "github.com/PMerdala/users-api/services/users"
	"github.com/PMerdala/users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"Implement me!")
}

func CreateUser(c *gin.Context){
	var user users.User
	if err:=c.ShouldBindJSON(&user); err!=nil{
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status,restErr)
		return
	}
	result,saveErr:=users_service.CreateUser(user)
	if saveErr!=nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}
	c.JSON(http.StatusCreated,result)
}
