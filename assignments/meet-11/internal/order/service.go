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



