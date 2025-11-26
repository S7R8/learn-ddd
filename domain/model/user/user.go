package user

type User struct {
	id   UserId
	name Name
	age  Age
}

func NewUserFromPrimitives(idStr string, nameStr string, ageInt int) (*User, error) {
	id, err := NewUserId(idStr)
	if err != nil {
		return nil, err
	}

	name, err := NewName(nameStr)
	if err != nil {
		return nil, err
	}

	age, err := NewAge(ageInt)
	if err != nil {
		return nil, err
	}

	return &User{id: id, name: name, age: age}, nil
}
