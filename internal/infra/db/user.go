package db

import (
	"context"
	"fmt"
)

func (q *Queries) SearchUsers(ctx context.Context, name string, filters Filters) ([]User, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), user_id, email, user_name, name, user_type, created_at, updated_at 
		FROM users 
		WHERE (name ILIKE $1 OR $1 = '') 
		ORDER BY %s %s, user_id ASC
        LIMIT $2 OFFSET $3`, filters.sortColumn(), filters.sortDirection())

	args := []any{name, filters.limit(), filters.offset()}
	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()
	totalRecords := 0

	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&totalRecords,
			&i.UserID,
			&i.Email,
			&i.UserName,
			&i.Name,
			&i.UserType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, Metadata{}, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return items, metadata, nil
}
