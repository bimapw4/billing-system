package presentations

import (
	"billing/internal/common"
	"time"
)

var (
	ErrPaymentNotFound     = common.Error("err payment not found")
	ErrPaymentAlreadyExist = common.Error("err payment already exist")
	ErrPaymentInstallment  = common.Error("err payment insufficient installments")
)

type Payment struct {
	ID        string     `db:"id" json:"id"`
	LoanID    string     `db:"loan_id" json:"loan_id"`
	Week      int        `db:"week" json:"week"`
	Paid      int        `db:"paid" json:"paid"`
	IsPaid    bool       `db:"is_paid" json:"is_paid"`
	PaidAt    *time.Time `db:"paid_at" json:"paid_at"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}

type RespPayment struct {
	Payment
	Deliquent       bool `json:"deliquent"`
	DeliquentAmount int  `json:"deliquent_amount"`
}

type PaymentStatus struct {
	IsDelinquent  bool  `json:"is_deliquent"`
	DeliquentWeek []int `json:"deliquent_week"`
	Outstanding   int   `json:"outstanding"`
}
