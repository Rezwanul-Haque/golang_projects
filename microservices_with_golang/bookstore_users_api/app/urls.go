package app

import (
	"../controllers/ping"
	"../controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	// GET: /users/1
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	// POST: /users
	router.POST("/users", users.CreateUser)
}