package user

import "errors"

type Age struct {
	value int
}

func NewAge(value int) (Age, error) {
	if value < 0 || value > 150 {
		return Age{}, errors.New("年齢は0〜150の範囲である必要があります")
	}
	return Age{value: value}, nil
}

func (a Age) Value() int {
	return a.value
}

// 値オブジェクトの等価性
func (a Age) Equals(other Age) bool {
	return a.value == other.value
}
