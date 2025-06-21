package repository

import (
	"context"
	"perpustakaan/internal/entity"

	"gorm.io/gorm"
)

type PenulisRepository struct {
	db *gorm.DB
}

func NewPenulisRepository(db *gorm.DB) *PenulisRepository {
	return &PenulisRepository{db: db}
}

func (r *PenulisRepository) Create(ctx context.Context, penulis *entity.Penulis) error {
	return r.db.WithContext(ctx).Create(penulis).Error
}

func (r *PenulisRepository) Update(ctx context.Context, penulis *entity.Penulis) error {
	return r.db.WithContext(ctx).Save(penulis).Error
}

func (r *PenulisRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Penulis{}, id).Error
}

func (r *PenulisRepository) GetByID(ctx context.Context, id int64) (*entity.Penulis, error) {
	var penulis entity.Penulis
	err := r.db.WithContext(ctx).First(&penulis, id).Error
	if err != nil {
		return nil, err
	}
	return &penulis, nil
}

func (r *PenulisRepository) GetAll(ctx context.Context) ([]entity.Penulis, error) {
	var penulis []entity.Penulis
	err := r.db.WithContext(ctx).Find(&penulis).Error
	if err != nil {
		return nil, err
	}
	return penulis, nil
}

func (r *PenulisRepository) GetByNama(ctx context.Context, nama string) (*entity.Penulis, error) {
	var penulis entity.Penulis
	err := r.db.WithContext(ctx).Where("nama_penulis = ?", nama).First(&penulis).Error
	if err != nil {
		return nil, err
	}
	return &penulis, nil
} 