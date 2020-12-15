package users

import (
	"fmt"
	"github.com/PMerdala/users-api/domain/users"
	usersService "github.com/PMerdala/users-api/services/users"
	"github.com/PMerdala/users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	userId, ok := getUserIdFromUrl(c)
	if !ok {
		return
	}
	user, getErr := usersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func GetByEmail(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, findErr := usersService.GetUserByEmail(user)
	if findErr != nil {
		c.JSON(findErr.Status, findErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := usersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Update(c *gin.Context) {
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
	result, saveErr := usersService.UpdateUser(isPartial, user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, ok := getUserIdFromUrl(c)
	if !ok {
		return
	}
	if deleteErr := usersService.DeleteUser(userId); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": fmt.Sprintf("delete user with id=%d", userId)})
}

func getUserIdFromUrl(c *gin.Context) (int64, bool) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(http.StatusBadRequest, err)
		return 0, false
	}
	return userId, true
}
