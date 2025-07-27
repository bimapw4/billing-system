package borrowers

import (
	"billing/internal/presentations"
	"context"
)

func (r *repo) Detail(ctx context.Context, id string) (*presentations.Borrower, error) {
	var (
		result = presentations.Borrower{}
	)

	query := `SELECT * FROM borrowers where id=:id`

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

func (r *repo) FindByUsername(ctx context.Context, username string) (*presentations.Borrower, error) {
	var (
		result = presentations.Borrower{}
	)

	query := `SELECT * FROM borrowers where name=:username`

	args := map[string]interface{}{
		"username": username,
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
