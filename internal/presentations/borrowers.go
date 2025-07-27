package presentations

import (
	"billing/internal/common"
	"time"
)

const (
	ErrBorrowersNotFound     = common.Error("err borrowers not found")
	ErrBorrowersAlreadyExist = common.Error("err borrowers already exist")
)

type Borrower struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Phone     string    `db:"phone" json:"phone"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
