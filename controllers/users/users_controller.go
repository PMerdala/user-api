package users

import (
	"github.com/PMerdala/users-api/domain/users"
	users_service "github.com/PMerdala/users-api/services/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"Implement me!")
}

func CreateUser(c *gin.Context){
	var user users.User
	if err:=c.ShouldBindJSON(&user); err!=nil{
		//TODO: Handle request error
		c.JSON(http.StatusInternalServerError,gin.H{"message":err.Error(),})
		return
	}
	result,saveErr:=users_service.CreateUser(user)
	if saveErr!=nil{
		//TODO: Handle user creation error
		c.JSON(http.StatusInternalServerError,gin.H{"message":saveErr.Error(),})
		return
	}
	c.JSON(http.StatusCreated,result)
}
