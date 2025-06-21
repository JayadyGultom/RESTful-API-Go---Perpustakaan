package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type BukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *BukuRepository {
	return &BukuRepository{db: db}
}

func (r *BukuRepository) Create(ctx context.Context, buku *entity.Buku) error {
	return r.db.WithContext(ctx).Create(buku).Error
}

func (r *BukuRepository) Update(ctx context.Context, buku *entity.Buku) error {
	return r.db.WithContext(ctx).Save(buku).Error
}

func (r *BukuRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Buku{}, id).Error
}

func (r *BukuRepository) GetByID(ctx context.Context, id int64) (*entity.Buku, error) {
	var buku entity.Buku
	err := r.db.WithContext(ctx).First(&buku, id).Error
	if err != nil {
		return nil, err
	}
	return &buku, nil
}

func (r *BukuRepository) GetAll(ctx context.Context) ([]entity.Buku, error) {
	var buku []entity.Buku
	err := r.db.WithContext(ctx).Find(&buku).Error
	if err != nil {
		return nil, err
	}
	return buku, nil
} 