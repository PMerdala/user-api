package users

import (
	"github.com/PMerdala/users-api/domain/users"
	"github.com/PMerdala/users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User)(*users.User,*errors.RestErr){
	var defaultUser users.User
	return &defaultUser,&errors.RestErr{Message: "Implement me!",Status:http.StatusNotImplemented, Error: "Implement me!"}
}