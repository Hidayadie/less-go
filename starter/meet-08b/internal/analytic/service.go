package analytic

import (
	"context"
	"encoding/json"
)

type Service struct{ Repository Repository }

// Saat ini hanya meneruskan ke repository. Tempat yang tepat untuk logika
// tambahan nanti, misalnya validasi filter ?country= atau pagination.
func (s Service) Customers(ctx context.Context) (json.RawMessage, error) {
	return s.Repository.Customers(ctx)
}

func (s Service) Products(ctx context.Context) (json.RawMessage, error) {
	return s.Repository.Products(ctx)
}
