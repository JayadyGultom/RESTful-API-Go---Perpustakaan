package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type PengembalianRepository struct {
	db *gorm.DB
}

func NewPengembalianRepository(db *gorm.DB) *PengembalianRepository {
	return &PengembalianRepository{db: db}
}

func (r *PengembalianRepository) Create(ctx context.Context, pengembalian *entity.Pengembalian) error {
	return r.db.WithContext(ctx).Create(pengembalian).Error
}

func (r *PengembalianRepository) Update(ctx context.Context, pengembalian *entity.Pengembalian) error {
	return r.db.WithContext(ctx).Save(pengembalian).Error
}

func (r *PengembalianRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Pengembalian{}, id).Error
}

func (r *PengembalianRepository) GetByID(ctx context.Context, id int64) (*entity.Pengembalian, error) {
	var pengembalian entity.Pengembalian
	err := r.db.WithContext(ctx).First(&pengembalian, id).Error
	if err != nil {
		return nil, err
	}
	return &pengembalian, nil
}

func (r *PengembalianRepository) GetAll(ctx context.Context) ([]entity.Pengembalian, error) {
	var pengembalians []entity.Pengembalian
	err := r.db.WithContext(ctx).Find(&pengembalians).Error
	if err != nil {
		return nil, err
	}
	return pengembalians, nil
} 