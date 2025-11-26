package user

import (
	"errors"
	"fmt"

	"learn-ddd/domain/model/user"

	"learn-ddd/domain/service"
)

type UserApplication struct {
	userService *service.UserService
}

func NewUserApplication(s *service.UserService) *UserApplication {
	return &UserApplication{userService: s}
}

func (a *UserApplication) CreateUser(user *user.User) error {
	if a.userService.DuplicateUser(user) {
		return errors.New("user already exists")
	}
	fmt.Println("OK")
	return nil
}
