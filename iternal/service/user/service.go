package user

import (
	"project/iternal/repository"
	"project/iternal/service"
)

type serv struct{
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{userRepository: userRepository}
}