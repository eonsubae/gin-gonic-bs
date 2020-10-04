package services

import (
	"github.com/eonsubae/gin-gonic-bs/domain/users"
	"github.com/eonsubae/gin-gonic-bs/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
