package handlers

import (
	"billing/internal/business"
	"billing/internal/handlers/borrowers"
	"billing/internal/handlers/loan"
)

type Handlers struct {
	Borrowers borrowers.Handler
	Loan      loan.Handler
}

func NewHandler(business business.Business) Handlers {
	return Handlers{
		Borrowers: borrowers.NewBorrowerHandler(business),
		Loan:      loan.NewLoanHandler(business),
	}
}
