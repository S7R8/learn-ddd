package main

import (
	userApp "learn-ddd/application/user"
	"learn-ddd/domain/model/user"
	"learn-ddd/domain/service"
)

func main() {
	app := userApp.NewUserApplication(service.NewUserService())
	app.CreateUser(&user.User{})
}
