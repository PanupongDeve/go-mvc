package app

import (
	healthzController "simple-service/controllers/healthz"
	usersController "simple-service/controllers/users"
)

func mapUrls() {
	router.GET("/healthz", healthzController.Healthz)
	router.POST("/users", usersController.CreateUser)
}
