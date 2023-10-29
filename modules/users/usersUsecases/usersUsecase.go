package usersUsecases

import (
	"github.com/forgetz/go-kawaii-shop/config"
	"github.com/forgetz/go-kawaii-shop/modules/users"
	"github.com/forgetz/go-kawaii-shop/modules/users/usersRepositories"
)

// interface, struct, constructor

type IUsersUsecase interface {
	InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error)
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

func (u *usersUsecase) InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error) {
	// hashing password
	if err := req.BcryptHashing(); err != nil {
		return nil, err
	}

	// insert user
	result, err := u.usersRepository.InsertUser(req, false)
	if err != nil {
		return nil, err
	}

	return result, nil
}
