package handlers

import (
	"billing/internal/business"
)

type Handlers struct {
}

func NewHandler(business business.Business) Handlers {
	return Handlers{}
}
