package repositories

import (
	"billing/internal/repositories/borrowers"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Borrowers borrowers.Borrowers
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Borrowers: borrowers.NewRepo(db),
	}
}
