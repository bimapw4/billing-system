package loan

import (
	"billing/internal/presentations"
	"billing/pkg/meta"
	"context"
)

type Loan interface {
	Create(ctx context.Context, input *presentations.Loan) error
	Detail(ctx context.Context, id string) (*presentations.Loan, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Loan, error)
}
