package main

import (
	"fmt"

	"learn-ddd/domain/model/user"
)

func main() {
	u, err := user.NewUserFromPrimitives("太郎", 30)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
