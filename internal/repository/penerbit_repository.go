package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type PenerbitRepository struct {
	db *gorm.DB
}

func NewPenerbitRepository(db *gorm.DB) *PenerbitRepository {
	return &PenerbitRepository{db: db}
}

func (r *PenerbitRepository) Create(ctx context.Context, penerbit *entity.Penerbit) error {
	return r.db.WithContext(ctx).Create(penerbit).Error
}

func (r *PenerbitRepository) Update(ctx context.Context, penerbit *entity.Penerbit) error {
	return r.db.WithContext(ctx).Save(penerbit).Error
}

func (r *PenerbitRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Penerbit{}, id).Error
}

func (r *PenerbitRepository) GetByID(ctx context.Context, id int64) (*entity.Penerbit, error) {
	var penerbit entity.Penerbit
	err := r.db.WithContext(ctx).First(&penerbit, id).Error
	if err != nil {
		return nil, err
	}
	return &penerbit, nil
}

func (r *PenerbitRepository) GetAll(ctx context.Context) ([]entity.Penerbit, error) {
	var penerbits []entity.Penerbit
	err := r.db.WithContext(ctx).Find(&penerbits).Error
	if err != nil {
		return nil, err
	}
	return penerbits, nil
}

func (r *PenerbitRepository) GetByNama(ctx context.Context, nama string) (*entity.Penerbit, error) {
	var penerbit entity.Penerbit
	err := r.db.WithContext(ctx).Where("nama_penerbit = ?", nama).First(&penerbit).Error
	if err != nil {
		return nil, err
	}
	return &penerbit, nil
}