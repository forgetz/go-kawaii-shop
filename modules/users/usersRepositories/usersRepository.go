package usersRepositories

import "github.com/jmoiron/sqlx"

// interface, struct, construtor

type IUsersRepository interface {
}

type usersRepository struct {
	db *sqlx.DB
}

func UsersRepository(db *sqlx.DB) IUsersRepository {
	return &usersRepository{
		db: db,
	}
}
