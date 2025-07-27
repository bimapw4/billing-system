package borrowers

import (
	"billing/internal/presentations"
	"billing/pkg/meta"
	"context"
)

type Borrowers interface {
	Create(ctx context.Context, input *presentations.Borrower) error
	Detail(ctx context.Context, id string) (*presentations.Borrower, error)
	FindByUsername(ctx context.Context, username string) (*presentations.Borrower, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Borrower, error)
}
