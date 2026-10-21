package order

import (
	"context"

	"gorm.io/gorm"
)


type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetOrderByID(ctx context.Context, id uint) (Order, error) {
	var order Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	return order, err
}

func (r *Repository) CreateOrder(ctx context.Context, order *Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *Repository) CreateOrderDetail(ctx context.Context, orderDetail *OrderDetail) error {
	return r.db.WithContext(ctx).Create(orderDetail).Error
}

func (r *Repository) GetOrderList(ctx context.Context, userID uint) ([]Order, error) {
	orders := []Order{}

	err := r.db.
		WithContext(ctx).
		Preload("user").
		Where("user_id = ?", userID).
		Order("id ASC").
		Limit(100).
		Find(&orders).Error
	return orders, err
}

func (r *Repository) GetOrderByIDWithUser(ctx context.Context, id uint, userID uint) (Order, error) {
	var order Order
	err := r.db.WithContext(ctx).Preload("user").Where("id = ? AND user_id = ?", id, userID).First(&order).Error
	return order, err
}

func (r *Repository) GetOrderDetailList(ctx context.Context) ([]OrderDetail, error) {
	order_details := []OrderDetail{}

	err := r.db.WithContext(ctx).Order("id ASC").Limit(100).Find(&order_details).Error
	return order_details, err
}
