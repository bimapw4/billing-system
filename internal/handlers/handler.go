package handlers

import (
	"billing/internal/business"
	"billing/internal/handlers/borrowers"
	"billing/internal/handlers/loan"
	"billing/internal/handlers/payment"
)

type Handlers struct {
	Borrowers borrowers.Handler
	Loan      loan.Handler
	Payment   payment.Handler
}

func NewHandler(business business.Business) Handlers {
	return Handlers{
		Borrowers: borrowers.NewBorrowerHandler(business),
		Loan:      loan.NewLoanHandler(business),
		Payment:   payment.NewLoanHandler(business),
	}
}
