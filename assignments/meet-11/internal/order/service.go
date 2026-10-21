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
/*
func (s *Service) Create(ctx context.Context, productID uint, quantity int, user auth.User) (Order, error) {
	if productID == 0 || quantity <= 0 || user.ID == 0 {
		return Order{}, ErrInvalidOrder
	}
	return s.repo.Create(ctx, productID, quantity, user.ID)
}
*/
func orderTotal(price, quantity int) (int, error) {
	if price < 0 || quantity <= 0 || (price > 0 && quantity > math.MaxInt/price) {
		return 0, ErrInvalidOrder
	}
	return price * quantity, nil
}

func (s *Service) List(ctx context.Context, userID uint) ([]Order, error) {
	return s.repo.List(ctx, userID)
}
func (s *Service) GetByID(ctx context.Context, id, userID uint) (*Order, error) {
	return s.repo.GetByID(ctx, id, userID)
}

/*
package order

import (
	"context"
	"errors"
	"gin-auth/internal/auth"
	//"gin-auth/internal/product"
)


type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(ctx context.Context, user auth.User, total float64) (Order, error) {
	order := Order{UserID: user.ID, Total: total}
	if err := s.repo.CreateOrder(ctx, &order); err != nil {
		return Order{}, err
	}
	order.User = user
	return order, nil


}
func (s *Service) CreateOrderDetail(ctx context.Context, user auth.User, orderID, productID uint, qty int, price float64) (OrderDetail, error) {
	// Validasi order milik user ini
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return OrderDetail{}, err
	}
	if order.UserID != user.ID {
		return OrderDetail{}, errors.New("order tidak ditemukan")
	}

	detail := OrderDetail{
		OrderID:   orderID,
		ProductID: productID,
		Qty:       qty,
		Price:     price,
		
	}
	if err := s.repo.CreateOrderDetail(ctx, &detail); err != nil {
		return OrderDetail{}, err
	}
	return detail, nil
}

func (s *Service) GetOrderList(ctx context.Context, userID uint) ([]Order, error) {
	return s.repo.GetOrderList(ctx, userID)
}

func (s *Service) GetOrderByID(ctx context.Context, id uint, userID uint) (Order, error) {
	return s.repo.GetOrderByIDWithUser(ctx, id, userID)
}

func (s *Service) GetOrderDetailList(ctx context.Context) ([]OrderDetail, error) {
	return s.repo.GetOrderDetailList(ctx)
}


*/
