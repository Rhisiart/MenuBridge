package types

import (
	"context"
	"database/sql"
)

type Encoded interface {
	Encode(data []byte, idx int, seq byte) (int, error)
	Type() byte
}

type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Table interface {
	Create(ctx context.Context, exec Executor) error
	Read(ctx context.Context, exec Executor) error
	ReadAll(ctx context.Context, exec Executor) ([]Table, error)
	Update(ctx context.Context, exec Executor) error
	Delete(ctx context.Context, exec Executor) error
}
