package order

import (
	"context"
	"errors"
	"gin-auth/internal/product"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }


func (r *Repository) Create(ctx context.Context, items []OrderItem, userID uint) (Order, error) {
	var order Order

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var grandTotal int

		// Kumpulkan details dulu, hitung total
		details := make([]OrderDetail, 0, len(items))
		for _, item := range items {
				var p product.Product
				if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
						First(&p, item.ProductID).Error; err != nil {
						return err
				}
				if p.Stock < item.Quantity {
						return ErrInsufficientStock
				}

				subtotal, err := orderTotal(p.Price, item.Quantity)
				if err != nil {
						return err
				}
				grandTotal += subtotal

				details = append(details, OrderDetail{
						ProductID: p.ID,
						Qty:       item.Quantity,
						Price:     p.Price,
				})

				if err := tx.Model(&p).
						Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
						return err
				}
		}

		// Buat order header
		order = Order{UserID: userID, Total: grandTotal}
		if err := tx.Create(&order).Error; err != nil {
				return err
		}

		// Isi OrderID ke semua details, lalu batch insert
		for i := range details {
				details[i].OrderID = order.ID
		}
		if err := tx.Create(&details).Error; err != nil {
				return err
		}

		return nil
	})
	if err != nil {
			return Order{}, err
	}

	// Preload di luar transaksi (seperti yang dibahas sebelumnya)
	if err := r.db.WithContext(ctx).
			Preload("User").
			Preload("Details.Product").
			First(&order, order.ID).Error; err != nil {
			return Order{}, err
	}

	return order, nil
}
/*
func (r *Repository) Create(ctx context.Context, productID uint, quantity int, userID uint) (Order, error) {
	var order Order
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var item product.Product
		// SELECT FOR UPDATE mengunci baris product sampai transaksi commit atau rollback.
		// Order bersamaan pada product yang sama harus menunggu, lalu membaca stok terbaru
		// agar stok yang sama tidak terjual dua kali.
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&item, productID).Error; err != nil {
			return err
		}

		if item.Stock < quantity {
			return ErrInsufficientStock
		}

		total, err := orderTotal(item.Price, quantity)
		if err != nil {
			return err
		}
		order = Order{UserID: userID, Total: total}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		detail := OrderDetail{OrderID: order.ID, ProductID: item.ID, Qty: quantity, Price: item.Price}
		if err := tx.Create(&detail).Error; err != nil {
			return err
		}

		if err := tx.Model(&item).Update("stock", gorm.Expr("stock - ?", quantity)).Error; err != nil {
			return err
		}

		return tx.Preload("User").Preload("Details.Product").First(&order, order.ID).Error
	})
	if err != nil {
		return Order{}, err
	}
	return order, nil
}
*/
func (r *Repository) List(ctx context.Context, userID uint) ([]Order, error) {
	orders := []Order{}
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Preload("User").Preload("Details.Product").Order("id DESC").Find(&orders).Error
	return orders, err
}

func (r *Repository) GetByID(ctx context.Context, id, userID uint) (*Order, error) {
	var order Order
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Preload("User").Preload("Details.Product").First(&order, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &order, nil
}
/*
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
*/
