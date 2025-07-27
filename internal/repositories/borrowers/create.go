package borrowers

import (
	"billing/internal/presentations"
	"context"
)

func (r *repo) Create(ctx context.Context, input *presentations.Borrower) error {

	query := `
    INSERT INTO borrowers (
        id, name, phone, created_at, updated_at
    ) VALUES (
        :id, :name, :phone, :created_at, :updated_at
    )`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         input.ID,
		"name":       input.Name,
		"phone":      input.Phone,
		"created_at": input.CreatedAt,
		"updated_at": input.UpdatedAt,
	})

	return r.translateError(err)
}
