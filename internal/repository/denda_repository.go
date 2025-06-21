package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type DendaRepository struct {
	db *gorm.DB
}

func NewDendaRepository(db *gorm.DB) *DendaRepository {
	return &DendaRepository{db: db}
}

func (r *DendaRepository) Create(ctx context.Context, denda *entity.Denda) error {
	return r.db.WithContext(ctx).Create(denda).Error
}

func (r *DendaRepository) Update(ctx context.Context, denda *entity.Denda) error {
	return r.db.WithContext(ctx).Save(denda).Error
}

func (r *DendaRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Denda{}, id).Error
}

func (r *DendaRepository) GetByID(ctx context.Context, id int64) (*entity.Denda, error) {
	var denda entity.Denda
	err := r.db.WithContext(ctx).First(&denda, id).Error
	if err != nil {
		return nil, err
	}
	return &denda, nil
}

func (r *DendaRepository) GetAll(ctx context.Context) ([]entity.Denda, error) {
	var denda []entity.Denda
	err := r.db.WithContext(ctx).Find(&denda).Error
	if err != nil {
		return nil, err
	}
	return denda, nil
} 