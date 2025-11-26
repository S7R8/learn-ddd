package service

import "learn-ddd/domain/model/user"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) DuplicateUser(user *user.User) bool {
	return false
}
