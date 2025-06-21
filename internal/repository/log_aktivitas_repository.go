package repository

import (
	"context"
	"perpustakaan/internal/entity"
	"gorm.io/gorm"
)

type LogAktivitasRepository struct {
	db *gorm.DB
}

func NewLogAktivitasRepository(db *gorm.DB) *LogAktivitasRepository {
	return &LogAktivitasRepository{db: db}
}

func (r *LogAktivitasRepository) Create(ctx context.Context, log *entity.LogAktivitas) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *LogAktivitasRepository) Update(ctx context.Context, log *entity.LogAktivitas) error {
	return r.db.WithContext(ctx).Save(log).Error
}

func (r *LogAktivitasRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.LogAktivitas{}, id).Error
}

func (r *LogAktivitasRepository) GetByID(ctx context.Context, id int64) (*entity.LogAktivitas, error) {
	var log entity.LogAktivitas
	err := r.db.WithContext(ctx).First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *LogAktivitasRepository) GetAll(ctx context.Context) ([]entity.LogAktivitas, error) {
	var logs []entity.LogAktivitas
	err := r.db.WithContext(ctx).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
} 