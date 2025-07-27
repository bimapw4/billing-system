package borrowers

import (
	"billing/internal/presentations"
	"billing/pkg/databasex"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Borrowers {
	return &repo{
		db: db,
	}
}

func (r *repo) translateError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return presentations.ErrBorrowersNotFound
	case databasex.ErrUniqueViolation:
		return presentations.ErrBorrowersAlreadyExist
	default:
		return err
	}
}
