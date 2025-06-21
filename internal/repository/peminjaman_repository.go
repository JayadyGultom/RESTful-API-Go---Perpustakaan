package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type PeminjamanRepository struct {
	db *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) *PeminjamanRepository {
	return &PeminjamanRepository{db: db}
}

func (r *PeminjamanRepository) Create(ctx context.Context, peminjaman *entity.Peminjaman) error {
	return r.db.WithContext(ctx).Create(peminjaman).Error
}

func (r *PeminjamanRepository) Update(ctx context.Context, peminjaman *entity.Peminjaman) error {
	return r.db.WithContext(ctx).Save(peminjaman).Error
}

func (r *PeminjamanRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Peminjaman{}, id).Error
}

func (r *PeminjamanRepository) GetByID(ctx context.Context, id int64) (*entity.Peminjaman, error) {
	var peminjaman entity.Peminjaman
	err := r.db.WithContext(ctx).First(&peminjaman, id).Error
	if err != nil {
		return nil, err
	}
	return &peminjaman, nil
}

func (r *PeminjamanRepository) GetAll(ctx context.Context) ([]entity.Peminjaman, error) {
	var peminjamans []entity.Peminjaman
	err := r.db.WithContext(ctx).Find(&peminjamans).Error
	if err != nil {
		return nil, err
	}
	return peminjamans, nil
} 