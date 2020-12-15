package app

import (
	"github.com/PMerdala/users-api/controllers/ping"
	"github.com/PMerdala/users-api/controllers/users"
)

func mapUrls() {
	route.GET("/ping", ping.Ping)
	route.GET("/users/:user_id", users.Get)
	route.GET("/search/users", users.GetByEmail)
	route.POST("/users", users.Create)
	route.PUT("/users/:user_id", users.Update)
	route.PATCH("/users/:user_id", users.Update)
	route.DELETE("/users/:user_id", users.Delete)
}
