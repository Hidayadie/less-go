package product

import (
	"context"
	"errors"
	"strings"

	"gin-auth/internal/auth"

	"gorm.io/gorm"
)

// ErrProductNotFound dipakai ketika product tidak ada atau bukan milik pengguna.
var ErrProductNotFound = errors.New("product tidak ditemukan")

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, user auth.User) (Product, error) {
	item := Product{Name: strings.TrimSpace(name), UserID: user.ID}
	if err := s.repo.Create(ctx, &item); err != nil {
		return Product{}, err
	}
	item.User = user
	return item, nil
}

func (s *Service) List(ctx context.Context, userID uint) ([]Product, error) {
	return s.repo.ListByUserID(ctx, userID)
}

// ListByID mengembalikan product milik pengguna yang sedang login.
func (s *Service) ListByID(ctx context.Context, productID uint, user auth.User) (Product, error) {
	return s.GetOwned(ctx, productID, user.ID)
}

// Put memperbarui nama product yang dimiliki pengguna lalu mengembalikan daftar product miliknya.
func (s *Service) Put(ctx context.Context, productID uint, item Product) ([]Product, error) {
	if _, err := s.GetOwned(ctx, productID, item.UserID); err != nil {
		return nil, err
	}
	if err := s.repo.UpdateByID(ctx, productID, item.Name); err != nil {
		return nil, err
	}
	return s.repo.ListByUserID(ctx, item.UserID)
}

// Delete menghapus product yang dimiliki pengguna.
func (s *Service) Delete(ctx context.Context, productID uint, user auth.User) error {
	if _, err := s.GetOwned(ctx, productID, user.ID); err != nil {
		return err
	}
	return s.repo.DeleteByID(ctx, productID)
}

// GetOwned mengembalikan product beserta pembuatnya bila product itu milik userID.
func (s *Service) GetOwned(ctx context.Context, productID uint, userID uint) (Product, error) {
	item, err := s.repo.GetByID(ctx, productID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Product{}, ErrProductNotFound
	}
	if err != nil {
		return Product{}, err
	}
	if item.UserID != userID {
		return Product{}, ErrProductNotFound
	}
	return item, nil
}
