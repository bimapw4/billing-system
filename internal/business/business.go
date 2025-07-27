package business

import (
	"billing/internal/business/borrowers"
	"billing/internal/repositories"
)

type Business struct {
	Borrowers borrowers.Borrower
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{
		Borrowers: borrowers.NewBusiness(repo),
	}
}
