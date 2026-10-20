package product

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

func (r *Repository) Create(ctx context.Context, item *Product) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *Repository) ListByUserID(ctx context.Context, userID uint) ([]Product, error) {
	items := []Product{}
	// ponytail: daftar maksimal 100 product terbaru; tambah pagination saat data melebihi batas.
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).
		Preload("User").Order("id DESC").Limit(100).Find(&items).Error
	return items, err
}

func (r *Repository) ListByID(ctx context.Context, productID uint) (Product, error) {
	var item Product
	err := r.db.WithContext(ctx).Preload("User").First(&item, productID).Error
	return item, err
}

func (r *Repository) GetByID(ctx context.Context, productID uint) (Product, error) {
	var item Product
	err := r.db.WithContext(ctx).Preload("User").First(&item, productID).Error
	return item, err
}

func (r *Repository) UpdateByID(ctx context.Context, productID uint, name string) error {
	return r.db.WithContext(ctx).Model(&Product{}).
		Where("id = ?", productID).
		Updates(map[string]any{"name": name}).Error
}

func (r *Repository) DeleteByID(ctx context.Context, productID uint) error {
	return r.db.WithContext(ctx).Delete(&Product{}, productID).Error
}

// Helper to check if product exists
func (r *Repository) Exists(ctx context.Context, productID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&Product{}).Where("id = ?", productID).Count(&count).Error
	return count > 0, err
}

