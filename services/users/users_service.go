package users

import (
	"github.com/PMerdala/users-api/domain/users"
	"github.com/PMerdala/users-api/utils/errors"
	"strings"
)

func CreateUser(user users.User)(*users.User,*errors.RestErr){
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email==""{
		return nil,errors.NewBadRequestError("invalid email address")
	}
	return nil,nil
}