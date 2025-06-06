package user

import (
	"project/internal/repository"
	"project/internal/service"
)

type serv struct{
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{userRepository: userRepository}
}