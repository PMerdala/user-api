package users

import "github.com/PMerdala/users-api/domain/users"

func CreateUser(user users.User)(*users.User,error){
	return &user,nil
}