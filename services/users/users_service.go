package users

import (
	"simple-service/domain/users"
	"simple-service/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
