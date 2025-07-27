package business

import (
	"billing/internal/business/borrowers"
	"billing/internal/business/loan"
	"billing/internal/business/payment"
	"billing/internal/repositories"
)

type Business struct {
	Borrowers borrowers.Borrower
	Loan      loan.Loan
	Payment   payment.Payment
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{
		Borrowers: borrowers.NewBusiness(repo),
		Loan:      loan.NewBusiness(repo),
		Payment:   payment.NewBusiness(repo),
	}
}
