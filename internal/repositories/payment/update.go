package payment

import (
	"billing/internal/presentations"
	"context"
	"time"
)

func (r *repo) UpdatePayment(ctx context.Context, payload presentations.Payment) error {
	query := `
	update payments set is_paid=:is_paid, paid=:paid, paid_at=:paid_at, updated_at=:updated_at where loan_id=:loan_id and week=:week
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"loan_id":    payload.LoanID,
		"week":       payload.Week,
		"is_paid":    payload.IsPaid,
		"paid":       payload.Paid,
		"paid_at":    payload.PaidAt,
		"updated_at": time.Now(),
	})

	return err
}
