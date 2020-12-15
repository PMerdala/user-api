package users

import (
	"github.com/PMerdala/users-api/domain/users"
	"github.com/PMerdala/users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetUserByEmail(searchUser users.User) (*users.User, *errors.RestErr) {
	if err := searchUser.CleanAndValidate(); err != nil {
		return nil, err
	}
	if err := searchUser.GetByEmail(); err != nil {
		return nil, err
	}
	return &searchUser, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.CleanAndValidate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if !isPartial || user.Email != "" {
		if err := user.CleanAndValidate(); err != nil {
			return nil, err
		}
	}
	if !isPartial || user.FirstName != "" {
		currentUser.FirstName = user.FirstName
	}
	if !isPartial || user.LastName != "" {
		currentUser.LastName = user.LastName
	}
	if !isPartial || user.Email != "" {
		currentUser.Email = user.Email
	}
	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

func PatchUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.CleanAndValidate(); err != nil {
		return nil, err
	}
	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email
	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}
