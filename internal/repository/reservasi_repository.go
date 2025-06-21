package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type ReservasiRepository struct {
	db *gorm.DB
}

func NewReservasiRepository(db *gorm.DB) *ReservasiRepository {
	return &ReservasiRepository{db: db}
}

func (r *ReservasiRepository) Create(ctx context.Context, reservasi *entity.Reservasi) error {
	return r.db.WithContext(ctx).Create(reservasi).Error
}

func (r *ReservasiRepository) Update(ctx context.Context, reservasi *entity.Reservasi) error {
	return r.db.WithContext(ctx).Save(reservasi).Error
}

func (r *ReservasiRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Reservasi{}, id).Error
}

func (r *ReservasiRepository) GetByID(ctx context.Context, id int64) (*entity.Reservasi, error) {
	var reservasi entity.Reservasi
	err := r.db.WithContext(ctx).First(&reservasi, id).Error
	if err != nil {
		return nil, err
	}
	return &reservasi, nil
}

func (r *ReservasiRepository) GetAll(ctx context.Context) ([]entity.Reservasi, error) {
	var reservasis []entity.Reservasi
	err := r.db.WithContext(ctx).Find(&reservasis).Error
	if err != nil {
		return nil, err
	}
	return reservasis, nil
} 