package users

import (
	"github.com/PMerdala/users-api/domain/users"
	"github.com/PMerdala/users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := users.Validate(&user); err != nil {
		return nil, err
	}
	return nil, nil
}
