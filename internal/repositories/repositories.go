package repositories

import (
	"billing/internal/repositories/borrowers"
	"billing/internal/repositories/loan"
	"billing/internal/repositories/payment"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Borrowers borrowers.Borrowers
	Loans     loan.Loan
	Payments  payment.Payment
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Borrowers: borrowers.NewRepo(db),
		Loans:     loan.NewRepo(db),
		Payments:  payment.NewRepo(db),
	}
}
