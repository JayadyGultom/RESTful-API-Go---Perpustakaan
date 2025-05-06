package repository

import (
	"perpustakaan/internal/entity"

	"gorm.io/gorm"
)

type BukuRepository interface {
	GetAllBuku() ([]entity.Buku, error)
	GetBukuByID(id int) (entity.Buku, error)
	CreateBuku(b entity.Buku) (entity.Buku, error)
	UpdateBuku(id int, b entity.Buku) (entity.Buku, error)
	DeleteBuku(id int) error
}

type bukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) BukuRepository {
	return &bukuRepository{db: db}
}

func (r *bukuRepository) GetAllBuku() ([]entity.Buku, error) {
	var buku []entity.Buku
	if err := r.db.Find(&buku).Error; err != nil {
		return nil, err
	}
	return buku, nil
}

func (r *bukuRepository) GetBukuByID(id int) (entity.Buku, error) {
	var buku entity.Buku
	if err := r.db.First(&buku, "id_buku = ?", id).Error; err != nil {
		return entity.Buku{}, err
	}
	return buku, nil
}

func (r *bukuRepository) CreateBuku(b entity.Buku) (entity.Buku, error) {
	if err := r.db.Create(&b).Error; err != nil {
		return entity.Buku{}, err
	}
	return r.GetBukuByID(b.IDBuku)
}

func (r *bukuRepository) UpdateBuku(id int, b entity.Buku) (entity.Buku, error) {
	b.IDBuku = id
	if err := r.db.Save(&b).Error; err != nil {
		return entity.Buku{}, err
	}
	return r.GetBukuByID(id)
}

func (r *bukuRepository) DeleteBuku(id int) error {
	return r.db.Delete(&entity.Buku{}, "id_buku = ?", id).Error
}
