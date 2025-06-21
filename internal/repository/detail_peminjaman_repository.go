package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type DetailPeminjamanRepository struct {
	db *gorm.DB
}

func NewDetailPeminjamanRepository(db *gorm.DB) *DetailPeminjamanRepository {
	return &DetailPeminjamanRepository{db: db}
}

func (r *DetailPeminjamanRepository) Create(ctx context.Context, detail *entity.DetailPeminjaman) error {
	return r.db.WithContext(ctx).Create(detail).Error
}

func (r *DetailPeminjamanRepository) Update(ctx context.Context, detail *entity.DetailPeminjaman) error {
	return r.db.WithContext(ctx).Save(detail).Error
}

func (r *DetailPeminjamanRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.DetailPeminjaman{}, id).Error
}

func (r *DetailPeminjamanRepository) GetByID(ctx context.Context, id int64) (*entity.DetailPeminjaman, error) {
	var detail entity.DetailPeminjaman
	err := r.db.WithContext(ctx).First(&detail, id).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *DetailPeminjamanRepository) GetAll(ctx context.Context) ([]entity.DetailPeminjaman, error) {
	var details []entity.DetailPeminjaman
	err := r.db.WithContext(ctx).Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
} 