package app

import (
	"github.com/PMerdala/users-api/datasources/mysql/users_db"
	"github.com/gin-gonic/gin"
)

var (
	route = gin.Default()
)

func StartApplication() {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	mapUrls()
	route.Run(":8080")
}
