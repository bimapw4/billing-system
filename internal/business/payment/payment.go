package payment

import (
	"billing/internal/common"
	"billing/internal/entity"
	"billing/internal/presentations"
	"billing/internal/repositories"
	"billing/pkg/meta"
	"context"
	"time"
)

type Payment interface {
	Payment(ctx context.Context, payload *entity.Payment) (*presentations.Payment, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Payment, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Payment {
	return &business{
		repo: repo,
	}
}

func (b *business) Payment(ctx context.Context, payload *entity.Payment) (*presentations.Payment, error) {

	loans, err := b.repo.Loans.Detail(ctx, payload.LoanID)
	if err != nil {
		return nil, err
	}

	if loans.BorrowerID != payload.BorrowerID {
		return nil, presentations.ErrLoanNotFound
	}

	installments := loans.Principal + (loans.Principal * loans.InterestRate / 100)

	if installments != payload.Payment {
		return nil, presentations.ErrPaymentInstallment
	}

	weeks := common.GetWeekNumber(loans.StartDate, payload.Date)

	err = b.repo.Payments.UpdatePayment(ctx, presentations.Payment{
		Week:      weeks,
		LoanID:    payload.LoanID,
		Paid:      payload.Payment,
		IsPaid:    true,
		PaidAt:    &payload.Date,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Payment, error) {
	list, err := b.repo.Payments.List(ctx, m)
	if err != nil {
		return nil, err
	}
	return list, nil
}
