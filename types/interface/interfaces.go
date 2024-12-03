package types

import (
	"context"
	"database/sql"
)

type Encoded interface {
	Encode(data []byte, idx int, seq byte) (int, error)
	Type() byte
}

type Table interface {
	Transaction(ctx context.Context, db *sql.DB) error
	Create(ctx context.Context, db *sql.DB) error
	Read(ctx context.Context, db *sql.DB) error
	ReadAll(ctx context.Context, db *sql.DB) ([]Table, error)
	Update(ctx context.Context, db *sql.DB) error
	Delete(ctx context.Context, db *sql.DB) error
}
