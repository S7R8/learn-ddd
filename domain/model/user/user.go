package user

type User struct {
	name Name
	age  Age
}

func NewUserFromPrimitives(nameStr string, ageInt int) (*User, error) {
	name, err := NewName(nameStr)
	if err != nil {
		return nil, err
	}

	age, err := NewAge(ageInt)
	if err != nil {
		return nil, err
	}

	return &User{name: name, age: age}, nil
}
