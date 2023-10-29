package middlewaresRepositories

import "github.com/jmoiron/sqlx"

// integate, struct, constuctor
type IMiddlewaresRepository interface {
}

type middlewaresRepository struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) IMiddlewaresRepository {
	return &middlewaresRepository{
		db: db,
	}
}
