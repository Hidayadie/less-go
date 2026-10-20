package category

import (
	"context"
	"errors"
	"strings"
)

var (
	// ErrCategoryNotFound dipakai ketika kategori tidak ada.
	ErrCategoryNotFound = errors.New("category tidak ditemukan")
	// ErrInvalidCategoryName dipakai ketika nama kategori kosong atau terlalu panjang.
	ErrInvalidCategoryName = errors.New("nama category tidak valid")
)

// maxCategoryName adalah batas panjang nama kategori sesuai kolom `name` di database.
const maxCategoryName = 100

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string) (Category, error) {
	category := Category{Name: strings.TrimSpace(name)}
	if category.Name == "" || len(category.Name) > maxCategoryName {
		return Category{}, ErrInvalidCategoryName
	}
	if err := s.repo.Create(ctx, &category); err != nil {
		return Category{}, err
	}
	return category, nil
}

func (s *Service) Get(ctx context.Context) ([]Category, error) {
	return s.repo.List(ctx)
}
