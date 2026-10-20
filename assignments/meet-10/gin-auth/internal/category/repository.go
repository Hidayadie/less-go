package category

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

func (r *Repository) Create(ctx context.Context, cat *Category) error {
	return r.db.WithContext(ctx).Create(cat).Error
}

func (r *Repository) List(ctx context.Context) ([]Category, error) {
	categories := []Category{}
	// ponytail: kembalikan maksimal 100 kategori terbaru; tambah pagination saat data melebihi batas.
	err := r.db.WithContext(ctx).Order("id ASC").Limit(100).Find(&categories).Error
	return categories, err
}
