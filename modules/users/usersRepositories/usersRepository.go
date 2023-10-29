package usersRepositories

import (
	"github.com/forgetz/go-kawaii-shop/modules/users"
	"github.com/forgetz/go-kawaii-shop/modules/users/usersPatterns"
	"github.com/jmoiron/sqlx"
)

// interface, struct, construtor

type IUsersRepository interface {
	InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error)
}

type usersRepository struct {
	db *sqlx.DB
}

func UsersRepository(db *sqlx.DB) IUsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error) {
	result := usersPatterns.InsertUser(r.db, req, isAdmin)

	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}

	// get result from insert user
	user, err := result.Result()
	if err != nil {
		return nil, err
	}

	return user, nil
}
