package analytic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type Repository struct{ DB *sql.DB }

// query membungkus SELECT agar PostgreSQL yang membentuk array JSON.
// Endpoint ini read-only, jadi tidak ada pemetaan error constraint (23xxx)
// seperti di package product.
func (r Repository) query(ctx context.Context, query string, args ...any) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query = "WITH result AS (" + query + ") SELECT COALESCE(json_agg(result), '[]'::json) FROM result"
	var output []byte
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&output); err != nil {
		return nil, err
	}
	if !json.Valid(output) {
		return nil, errors.New("PostgreSQL menghasilkan JSON tidak valid")
	}
	return output, nil
}

func (r Repository) Customers(ctx context.Context) (json.RawMessage, error) {
	return r.query(ctx, customerAnalysis)
}

func (r Repository) Products(ctx context.Context) (json.RawMessage, error) {
	return r.query(ctx, productAnalysis)
}
