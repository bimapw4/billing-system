package loan

import (
	"billing/internal/presentations"
	"billing/pkg/meta"
	"context"
	"fmt"
	"strings"
)

func (r *repo) List(ctx context.Context, m *meta.Params) ([]presentations.Loan, error) {
	var (
		result = []presentations.Loan{}
	)

	q, err := meta.Parse(m)
	if err != nil {
		return nil, err
	}
	query := `SELECT * FROM loans ORDER BY created_at DESC OFFSET :offset LIMIT :limit`

	query = strings.Replace(
		query,
		" ORDER BY created_at DESC ",
		fmt.Sprintf(" ORDER BY %s %s ", q.OrderBy, q.OrderDirection),
		1,
	)

	if m.SearchBy != "" {
		query = strings.ReplaceAll(query, "1=1", fmt.Sprintf("%v='%v'", m.SearchBy, m.Search))
	}

	args := map[string]interface{}{
		"offset": q.Offset,
		"limit":  q.Limit,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.SelectContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}

	total, err := r.count(ctx)
	if err != nil {
		return nil, r.translateError(err)
	}

	m.TotalItems = total

	return result, nil
}

func (r *repo) count(ctx context.Context) (int, error) {

	result := 0

	query := `SELECT count(*) FROM loans`

	args := map[string]interface{}{}

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
