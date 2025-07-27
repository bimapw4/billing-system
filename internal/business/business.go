package business

import (
	"billing/internal/business/borrowers"
	"billing/internal/business/loan"
	"billing/internal/repositories"
)

type Business struct {
	Borrowers borrowers.Borrower
	Loan      loan.Loan
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{
		Borrowers: borrowers.NewBusiness(repo),
		Loan:      loan.NewBusiness(repo),
	}
}
