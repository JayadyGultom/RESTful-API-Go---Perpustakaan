package repository

import (
	"context"
	"perpustakaan/internal/entity"

	"gorm.io/gorm"
)

type KategoriBukuRepository struct {
	db *gorm.DB
}

func NewKategoriBukuRepository(db *gorm.DB) *KategoriBukuRepository {
	return &KategoriBukuRepository{db: db}
}

func (r *KategoriBukuRepository) Create(ctx context.Context, kategori *entity.KategoriBuku) error {
	return r.db.WithContext(ctx).Create(kategori).Error
}

func (r *KategoriBukuRepository) Update(ctx context.Context, kategori *entity.KategoriBuku) error {
	return r.db.WithContext(ctx).Save(kategori).Error
}

func (r *KategoriBukuRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.KategoriBuku{}, id).Error
}

func (r *KategoriBukuRepository) GetByID(ctx context.Context, id int64) (*entity.KategoriBuku, error) {
	var kategori entity.KategoriBuku
	err := r.db.WithContext(ctx).First(&kategori, id).Error
	if err != nil {
		return nil, err
	}
	return &kategori, nil
}

func (r *KategoriBukuRepository) GetAll(ctx context.Context) ([]entity.KategoriBuku, error) {
	var kategoris []entity.KategoriBuku
	err := r.db.WithContext(ctx).Find(&kategoris).Error
	if err != nil {
		return nil, err
	}
	return kategoris, nil
}

func (r *KategoriBukuRepository) GetByNama(ctx context.Context, nama string) (*entity.KategoriBuku, error) {
	var kategori entity.KategoriBuku
	err := r.db.WithContext(ctx).Where("nama_kategori = ?", nama).First(&kategori).Error
	if err != nil {
		return nil, err
	}
	return &kategori, nil
} 