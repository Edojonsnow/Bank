// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2 
)RETURNING id, account_id, amount, created_at
`

func (q *Queries) CreateEntry(ctx context.Context, accountID int64, amount int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, accountID, amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, id)
	return err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_id, amount, created_at FROM entries
ORDER BY id
LIMIT $1 OFFSET $2
`

func (q *Queries) ListEntries(ctx context.Context, limit int32, offset int32) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntry = `-- name: UpdateEntry :one
UPDATE entries
SET amount = $2
WHERE id = $1
RETURNING id, account_id, amount, created_at
`

func (q *Queries) UpdateEntry(ctx context.Context, iD int64, amount int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, updateEntry, iD, amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
