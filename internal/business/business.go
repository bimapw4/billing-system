package business

import (
	"billing/internal/repositories"
)

type Business struct {
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{}
}
