package loan

import (
	"billing/internal/presentations"
	"context"

	"github.com/google/uuid"
)

func (r *repo) Create(ctx context.Context, input *presentations.Loan) error {

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `
    INSERT INTO loans (
        id, borrower_id, principal, interest_rate, total_weeks, start_date, created_at, updated_at
    ) VALUES (
        :id, :borrower_id, :principal, :interest_rate, :total_weeks, :start_date, :created_at, :updated_at
    )`

	_, err = tx.NamedExecContext(ctx, query, map[string]interface{}{
		"id":            input.ID,
		"borrower_id":   input.BorrowerID,
		"principal":     input.Principal,
		"interest_rate": input.InterestRate,
		"total_weeks":   input.TotalWeeks,
		"start_date":    input.StartDate,
		"created_at":    input.CreatedAt,
		"updated_at":    input.UpdatedAt,
	})
	if err != nil {
		return r.translateError(err)
	}

	paymentInsertQuery := `INSERT INTO payments (id, loan_id, week, paid, is_paid, paid_at) VALUES ($1, $2, $3, $4, false, NULL)`
	stmt, err := tx.PrepareContext(ctx, paymentInsertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for week := 1; week <= input.TotalWeeks; week++ {
		_, err := stmt.ExecContext(ctx, uuid.NewString(), input.ID, week, input.WeeklyPayment)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return r.translateError(err)
}
