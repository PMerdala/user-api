package users

import (
	"encoding/json"
	"github.com/PMerdala/users-api/domain/users"
	users_service "github.com/PMerdala/users-api/services/users"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"Implement me!")
}

func CreateUser(c *gin.Context){
	var user users.User
	bytes,err:=ioutil.ReadAll(c.Request.Body)
	if err!=nil{
		//TODO: Handle error
		c.JSON(http.StatusInternalServerError,gin.H{"message":err.Error(),})
		return
	}
	if err:= json.Unmarshal(bytes,&user); err!=nil{
		//TODO: Handle json error
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
