package users

import (
	"../../utils/errors"
	"fmt"
	"time"
)

var(
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	currentUser := userDB[user.Id]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("User email %s already registered.", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists.", user.Id))
	}

	now := time.Now()
	user.DateCreated = now.Format("2006-02-06T21:44:05Z")

	userDB[user.Id] = user
	return nil
}