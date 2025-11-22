package user

import "errors"

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if value == "" {
		return Name{}, errors.New("名前は空にできません")
	}
	if len(value) > 50 {
		return Name{}, errors.New("名前は50文字以内である必要があります")
	}
	return Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}

func (n Name) Equals(other Name) bool {
	return n.value == other.value
}
