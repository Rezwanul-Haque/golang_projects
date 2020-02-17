package services

import (
	"../domain/users"
	"../utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return nil, nil
	return &user, nil
	return nil, &errors.RestError{
		Status: http.StatusInternalServerError,
	}
}