package payment

import (
	"billing/internal/presentations"
	"billing/pkg/meta"
	"context"
)

type Payment interface {
	UpdatePayment(ctx context.Context, payload presentations.Payment) error
	Detail(ctx context.Context, id string) (*presentations.Payment, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Payment, error)
	FindByWeeksAndLoanID(ctx context.Context, week int, loanID string) (*presentations.Payment, error)
	CountIsPaidFalse(ctx context.Context, loanID string) (int, error)
}
