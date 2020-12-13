package app

import (
	"github.com/PMerdala/users-api/controllers/ping"
	"github.com/PMerdala/users-api/controllers/users"
)

func mapUrls() {
	route.GET("/ping", ping.Ping)
	route.GET("/users/:user_id", users.GetUser)
	route.GET("/search/users", users.GetUserByEmail)
	route.POST("/users", users.CreateUser)
}
