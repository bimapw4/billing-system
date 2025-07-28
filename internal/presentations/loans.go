package presentations

import (
	"billing/internal/common"
	"time"
)

const (
	ErrLoanNotFound     = common.Error("err loan not found")
	ErrLoanAlreadyExist = common.Error("err loan already exist")
)

type Loan struct {
	ID            string    `db:"id" json:"id"`
	BorrowerID    string    `db:"borrower_id" json:"borrower_id"`
	Principal     int       `db:"principal" json:"principal"`
	InterestRate  int       `db:"interest_rate" json:"interest_rate"`
	TotalWeeks    int       `db:"total_weeks" json:"total_weeks"`
	StartDate     time.Time `db:"start_date" json:"start_date"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	WeeklyPayment int       `json:"weekly_payment,omitempty"`
}
