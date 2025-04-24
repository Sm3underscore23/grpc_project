package user

import (
	"project/iternal/service"
	desc "project/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewUserImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
