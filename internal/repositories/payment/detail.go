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

func (r *repo) FindByWeeksAndLoanID(ctx context.Context, week int, loanID string) (*presentations.Payment, error) {
	var (
		result = presentations.Payment{}
	)

	query := `SELECT * FROM payments where week=:week and loan_id=:loan_id`

	args := map[string]interface{}{
		"week":    week,
		"loan_id": loanID,
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

func (r *repo) CountIsPaidFalse(ctx context.Context, loanID string) (int, error) {
	var (
		result int
	)

	query := `SELECT  COUNT(*) FROM payments where loan_id=:loan_id and is_paid=false`

	args := map[string]interface{}{
		"loan_id": loanID,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, r.translateError(err)
	}

	err = stmt.GetContext(ctx, &result, args)
	if err != nil {
		return 0, r.translateError(err)
	}

	return result, nil
}
