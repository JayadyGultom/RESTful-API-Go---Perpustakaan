package repository

import (
	"context"
	"perpustakaan/internal/entity"

	"gorm.io/gorm"
)

type PetugasRepository struct {
	db *gorm.DB
}

func NewPetugasRepository(db *gorm.DB) *PetugasRepository {
	return &PetugasRepository{db: db}
}

func (r *PetugasRepository) Create(ctx context.Context, petugas *entity.Petugas) error {
	return r.db.WithContext(ctx).Create(petugas).Error
}

func (r *PetugasRepository) Update(ctx context.Context, petugas *entity.Petugas) error {
	return r.db.WithContext(ctx).Save(petugas).Error
}

func (r *PetugasRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Petugas{}, id).Error
}

func (r *PetugasRepository) GetByID(ctx context.Context, id int64) (*entity.Petugas, error) {
	var petugas entity.Petugas
	err := r.db.WithContext(ctx).First(&petugas, id).Error
	if err != nil {
		return nil, err
	}
	return &petugas, nil
}

func (r *PetugasRepository) GetAll(ctx context.Context) ([]entity.Petugas, error) {
	var petugas []entity.Petugas
	err := r.db.WithContext(ctx).Find(&petugas).Error
	if err != nil {
		return nil, err
	}
	return petugas, nil
}

func (r *PetugasRepository) GetByUsername(ctx context.Context, username string) (*entity.Petugas, error) {
	var petugas entity.Petugas
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&petugas).Error
	if err != nil {
		return nil, err
	}
	return &petugas, nil
} 