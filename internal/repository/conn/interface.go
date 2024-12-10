package conn

import (
	"context"
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"io"
)

type SQLConnection interface {
	io.Closer
	Exec(ctx context.Context, query string, args ...interface{}) error
	Query(ctx context.Context, query string, args ...interface{}) (*pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *pgx.Row
}

type repository struct {
	conn *pgx.Conn
}

func (r *repository) Close() error {
	return r.conn.Close()
}

func NewSQLRepository(ctx context.Context, config pgx.ConnConfig) (SQLConnection, error) {
	conn, err := pgx.Connect(config)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}
	return &repository{conn: conn}, nil
}

func (r *repository) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := r.conn.ExecEx(ctx, query, nil, args...)
	return err
}

func (r *repository) Query(ctx context.Context, query string, args ...interface{}) (*pgx.Rows, error) {
	rows, err := r.conn.QueryEx(ctx, query, nil, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *repository) QueryRow(ctx context.Context, query string, args ...interface{}) *pgx.Row {
	row := r.conn.QueryRowEx(ctx, query, nil, args...)
	return row
}
