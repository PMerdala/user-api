package users

import (
	"github.com/PMerdala/users-api/domain/users"
	"github.com/PMerdala/users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.CleanAndValidate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
