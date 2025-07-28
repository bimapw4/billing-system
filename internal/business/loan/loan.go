package loan

import (
	"billing/internal/entity"
	"billing/internal/presentations"
	"billing/internal/repositories"
	"billing/pkg/meta"
	"context"
	"time"

	"github.com/google/uuid"
)

type Loan interface {
	Create(ctx context.Context, payload *entity.Loan) (*presentations.Loan, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Loan, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Loan {
	return &business{
		repo: repo,
	}
}

func (b *business) Create(ctx context.Context, payload *entity.Loan) (*presentations.Loan, error) {

	data := presentations.Loan{
		ID:           uuid.NewString(),
		BorrowerID:   payload.BorrowerID,
		Principal:    payload.Principal,
		InterestRate: payload.InterestRate,
		TotalWeeks:   payload.TotalWeeks,
		StartDate:    payload.StartDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	totalToPay := data.Principal + int((float64(data.Principal) * (float64(data.InterestRate) / 100)))

	weeklyPayment := totalToPay / data.TotalWeeks

	data.WeeklyPayment = weeklyPayment

	err := b.repo.Loans.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Loan, error) {
	list, err := b.repo.Loans.List(ctx, m)
	if err != nil {
		return nil, err
	}
	return list, nil
}
