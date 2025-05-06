package repository

import (
	"perpustakaan/internal/entity"

	"gorm.io/gorm"
)

type KategoriBukuRepository interface {
	GetAllKategoriBuku() ([]entity.KategoriBuku, error)
	GetKategoriBukuByID(id int) (entity.KategoriBuku, error)
	CreateKategoriBuku(k entity.KategoriBuku) (entity.KategoriBuku, error)
	UpdateKategoriBuku(id int, k entity.KategoriBuku) (entity.KategoriBuku, error)
	DeleteKategoriBuku(id int) error
}

type kategoriBukuRepository struct {
	db *gorm.DB
}

func NewKategoriBukuRepository(db *gorm.DB) KategoriBukuRepository {
	return &kategoriBukuRepository{db: db}
}

func (r *kategoriBukuRepository) GetAllKategoriBuku() ([]entity.KategoriBuku, error) {
	var kategori []entity.KategoriBuku
	if err := r.db.Find(&kategori).Error; err != nil {
		return nil, err
	}
	return kategori, nil
}

func (r *kategoriBukuRepository) GetKategoriBukuByID(id int) (entity.KategoriBuku, error) {
	var kategori entity.KategoriBuku
	if err := r.db.First(&kategori, "id_kategori = ?", id).Error; err != nil {
		return entity.KategoriBuku{}, err
	}
	return kategori, nil
}

func (r *kategoriBukuRepository) CreateKategoriBuku(k entity.KategoriBuku) (entity.KategoriBuku, error) {
	if err := r.db.Create(&k).Error; err != nil {
		return entity.KategoriBuku{}, err
	}
	return r.GetKategoriBukuByID(k.IDKategori)
}

func (r *kategoriBukuRepository) UpdateKategoriBuku(id int, k entity.KategoriBuku) (entity.KategoriBuku, error) {
	k.IDKategori = id
	if err := r.db.Save(&k).Error; err != nil {
		return entity.KategoriBuku{}, err
	}
	return r.GetKategoriBukuByID(id)
}

func (r *kategoriBukuRepository) DeleteKategoriBuku(id int) error {
	return r.db.Delete(&entity.KategoriBuku{}, "id_kategori = ?", id).Error
}
