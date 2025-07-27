package handlers

import (
	"billing/internal/business"
	"billing/internal/handlers/borrowers"
)

type Handlers struct {
	Borrowers borrowers.Handler
}

func NewHandler(business business.Business) Handlers {
	return Handlers{
		Borrowers: borrowers.NewAdminHandler(business),
	}
}
