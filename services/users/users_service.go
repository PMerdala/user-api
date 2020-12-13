package users

import (
	"github.com/PMerdala/users-api/domain"
	"github.com/PMerdala/users-api/domain/users"
)

func CreateUser(user users.User)(*users.User,*domain.Error){
	return &user,nil
}