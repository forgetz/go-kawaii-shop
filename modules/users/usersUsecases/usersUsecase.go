package usersUsecases

import (
	"github.com/forgetz/go-kawaii-shop/config"
	"github.com/forgetz/go-kawaii-shop/modules/users/usersRepositories"
)

// interface, struct, constructor

type IUsersUsecase interface {
}

type usersUsecase struct {
	cfg             config.IConfig
	usersRepository usersRepositories.IUsersRepository
}

func UsersUsecase(cfg config.IConfig, usersRepository usersRepositories.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		cfg:             cfg,
		usersRepository: usersRepository,
	}
}
