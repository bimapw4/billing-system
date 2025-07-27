package loan

import (
	"billing/internal/presentations"
	"context"
)

func (r *repo) Detail(ctx context.Context, id string) (*presentations.Loan, error) {
	var (
		result = presentations.Loan{}
	)

	query := `SELECT * FROM loans where id=:id`

	args := map[string]interface{}{
		"id": id,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.GetContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}

	return &result, nil
}
