package user

import "errors"

type UserId struct {
	value string
}

func NewUserId(value string) (UserId, error) {
	if value == "" {
		return UserId{}, errors.New("UserIdは空にできません")
	}
	return UserId{value: value}, nil
}

func (id UserId) Value() string {
	return id.value
}

func (id UserId) Equals(other UserId) bool {
	return id.value == other.value
}
