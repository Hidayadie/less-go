package order

import (
	"context"
	"errors"
	"math"

	"gin-auth/internal/auth"
)

var (
	ErrInvalidOrder      = errors.New("product_id dan quantity harus lebih dari 0")
	ErrInsufficientStock = errors.New("stok product tidak cukup")
)

type OrderItem struct {
	ProductID uint
	Quantity  int
}

type Service struct{ repo *Repository }

func NewService(repo *Repository) *Service { return &Service{repo: repo} }

func (s *Service) Create(ctx context.Context, items []OrderItem, user auth.User) (Order, error) {
	if len(items) == 0 || user.ID == 0 {
		return Order{}, ErrInvalidOrder
	}
	for _, item := range items {
		if item.ProductID == 0 || item.Quantity <= 0 {
			return Order{}, ErrInvalidOrder
		}
	}
	return s.repo.Create(ctx, items, user.ID)
}

func orderTotal(price, quantity int) (int, error) {
	if price < 0 || quantity <= 0 || (price > 0 && quantity > math.MaxInt/price) {
		return 0, ErrInvalidOrder
	}
	return price * quantity, nil
}

func taxSubtotal(total int) (int, error) {
	if total < 0 {
		return 0, ErrInvalidOrder
	}
	tax := int(float64(total) * 0.11)
	subtotal := total + tax
	return subtotal, nil
}

func (s *Service) List(ctx context.Context, userID uint) ([]Order, error) {
	return s.repo.List(ctx, userID)
}
func (s *Service) GetByID(ctx context.Context, id, userID uint) (*Order, error) {
	return s.repo.GetByID(ctx, id, userID)
}

