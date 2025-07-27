package repositories

import (
	"billing/internal/repositories/borrowers"
	"billing/internal/repositories/loan"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Borrowers borrowers.Borrowers
	Loans     loan.Loan
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Borrowers: borrowers.NewRepo(db),
		Loans:     loan.NewRepo(db),
	}
}
