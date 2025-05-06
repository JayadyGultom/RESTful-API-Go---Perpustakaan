package repository

import (
	"perpustakaan/internal/entity"

	"gorm.io/gorm"
)

type DetailPeminjamanRepository interface {
	GetAllDetailPeminjaman() ([]entity.DetailPeminjaman, error)
	GetDetailPeminjamanByID(id int) (entity.DetailPeminjaman, error)
	CreateDetailPeminjaman(d entity.DetailPeminjaman) (entity.DetailPeminjaman, error)
	UpdateDetailPeminjaman(id int, d entity.DetailPeminjaman) (entity.DetailPeminjaman, error)
	DeleteDetailPeminjaman(id int) error
}

type detailPeminjamanRepository struct {
	db *gorm.DB
}

func NewDetailPeminjamanRepository(db *gorm.DB) DetailPeminjamanRepository {
	return &detailPeminjamanRepository{db: db}
}

func (r *detailPeminjamanRepository) GetAllDetailPeminjaman() ([]entity.DetailPeminjaman, error) {
	var detail []entity.DetailPeminjaman
	if err := r.db.Find(&detail).Error; err != nil {
		return nil, err
	}
	return detail, nil
}

func (r *detailPeminjamanRepository) GetDetailPeminjamanByID(id int) (entity.DetailPeminjaman, error) {
	var detail entity.DetailPeminjaman
	if err := r.db.First(&detail, "id_detail = ?", id).Error; err != nil {
		return entity.DetailPeminjaman{}, err
	}
	return detail, nil
}

func (r *detailPeminjamanRepository) CreateDetailPeminjaman(d entity.DetailPeminjaman) (entity.DetailPeminjaman, error) {
	if err := r.db.Create(&d).Error; err != nil {
		return entity.DetailPeminjaman{}, err
	}
	return r.GetDetailPeminjamanByID(d.IDDetail)
}

func (r *detailPeminjamanRepository) UpdateDetailPeminjaman(id int, d entity.DetailPeminjaman) (entity.DetailPeminjaman, error) {
	d.IDDetail = id
	if err := r.db.Save(&d).Error; err != nil {
		return entity.DetailPeminjaman{}, err
	}
	return r.GetDetailPeminjamanByID(id)
}

func (r *detailPeminjamanRepository) DeleteDetailPeminjaman(id int) error {
	return r.db.Delete(&entity.DetailPeminjaman{}, "id_detail = ?", id).Error
}
