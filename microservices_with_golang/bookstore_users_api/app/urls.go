package app

import (
	"../controllers/ping"
	"../controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	// POST: /users
	router.POST("/users", users.Create)
	// GET: /users/1
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}