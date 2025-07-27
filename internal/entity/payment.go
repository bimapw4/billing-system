package entity

import "time"

type Payment struct {
	BorrowerID string    `json:"borrower_id"`
	LoanID     string    `json:"loan_id"`
	Payment    int       `json:"payment"`
	Date       time.Time `json:"date"`
}
