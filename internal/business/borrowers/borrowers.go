package borrowers

import (
	"billing/internal/entity"
	"billing/internal/presentations"
	"billing/internal/repositories"
	"billing/pkg/meta"
	"context"
	"time"

	"github.com/google/uuid"
)

type Borrower interface {
	Create(ctx context.Context, payload *entity.Borrower) (*presentations.Borrower, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Borrower, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Borrower {
	return &business{
		repo: repo,
	}
}

func (b *business) Create(ctx context.Context, payload *entity.Borrower) (*presentations.Borrower, error) {

	ext, _ := b.repo.Borrowers.FindByUsername(ctx, payload.Username)
	if ext != nil {
		return nil, presentations.ErrBorrowersAlreadyExist
	}

	data := presentations.Borrower{
		ID:        uuid.NewString(),
		Name:      payload.Username,
		Phone:     payload.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := b.repo.Borrowers.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Borrower, error) {
	list, err := b.repo.Borrowers.List(ctx, m)
	if err != nil {
		return nil, err
	}
	return list, nil
}
