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

		substot, err := taxSubtotal(grandTotal)
		if err != nil {
 			return err
		}
		
		order = Order{
			UserID: userID, 
			Total: grandTotal,
			SubTotal: substot,
		}
		if err := tx.Create(&order).Error; err != nil {
				return err
		}
		

		// isi OrderID 
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

	// preload 
	if err := r.db.WithContext(ctx).
			Preload("User").
			Preload("Details.Product").
			First(&order, order.ID).Error; err != nil {
			return Order{}, err
	}

	return order, nil
}

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
