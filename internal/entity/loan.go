package entity

import "time"

type Loan struct {
	BorrowerID   string    `json:"borrower_id"`
	Principal    int       `json:"principal"`
	InterestRate int       `json:"interest_rate"`
	TotalWeeks   int       `json:"total_weeks"`
	StartDate    time.Time `json:"start_date"`
}
