package usersHandlers

import (
	"github.com/forgetz/go-kawaii-shop/config"
	"github.com/forgetz/go-kawaii-shop/modules/users/usersUsecases"
)

// interface, struct, constructor

type IUsersHandler interface {
}

type usersHandler struct {
	cfg          config.IConfig
	usersUsecase usersUsecases.IUsersUsecase
}

func UsersHandler(cfg config.IConfig, usersUsecase usersUsecases.IUsersUsecase) IUsersHandler {
	return &usersHandler{
		cfg:          cfg,
		usersUsecase: usersUsecase,
	}
}
