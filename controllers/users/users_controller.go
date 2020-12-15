package users

import (
	"github.com/PMerdala/users-api/domain/users"
	users_service "github.com/PMerdala/users-api/services/users"
	"github.com/PMerdala/users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, ok := getUserIdFromUrl(c)
	if !ok {
		return
	}
	user, getErr := users_service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func GetUserByEmail(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, findErr := users_service.GetUserByEmail(user)
	if findErr != nil {
		c.JSON(findErr.Status, findErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := users_service.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	userId, ok := getUserIdFromUrl(c)
	if !ok {
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, saveErr := users_service.UpdateUser(isPartial, user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, ok := getUserIdFromUrl(c)
	if !ok {
		return
	}
	result, deleteErr := users_service.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func getUserIdFromUrl(c *gin.Context) (userId int64, status bool) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(http.StatusBadRequest, err)
		return 0, false
	}
	return userId, true
}
