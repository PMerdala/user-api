package users

import (
	"fmt"
	"github.com/PMerdala/users-api/utils/date_utils"
	"github.com/PMerdala/users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User not found by id: %d", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	if user.Id == 0 {
		user.Id = int64(len(usersDB)) + 1
	} else {
		current := usersDB[user.Id]
		if current != nil {
			if current.Email == user.Email {
				return errors.NewBadRequestError(fmt.Sprintf("User already exists with email: %s", user.Email))
			}
			return errors.NewBadRequestError(fmt.Sprintf("User already exists with id: %d", user.Id))
		}
	}
	foundUserId := int64(0)
	for userId,currentUser:=range usersDB{
		if currentUser.Email == user.Email{
			foundUserId=userId
			break
		}
	}
	if foundUserId != 0{
		return errors.NewBadRequestError(fmt.Sprintf("User already exists with email: %s", user.Email))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
