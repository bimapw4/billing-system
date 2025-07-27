package payment

import (
	"billing/internal/presentations"
	"context"
)

func (r *repo) Detail(ctx context.Context, id string) (*presentations.Payment, error) {
	var (
		result = presentations.Payment{}
	)

	query := `SELECT * FROM payments where id=:id`

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
