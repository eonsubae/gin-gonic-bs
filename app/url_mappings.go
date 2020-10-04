package app

import (
	"github.com/eonsubae/gin-gonic-bs/controllers/ping"
	"github.com/eonsubae/gin-gonic-bs/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
