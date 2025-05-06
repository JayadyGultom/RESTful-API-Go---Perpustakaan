package repository

import (
	"perpustakaan/internal/entity"

	"gorm.io/gorm"
)

type PeminjamanRepository interface {
	GetAllPeminjaman() ([]entity.Peminjaman, error)
	GetPeminjamanByID(id int) (entity.Peminjaman, error)
	CreatePeminjaman(p entity.Peminjaman) (entity.Peminjaman, error)
	UpdatePeminjaman(id int, p entity.Peminjaman) (entity.Peminjaman, error)
	DeletePeminjaman(id int) error
}

type peminjamanRepository struct {
	db *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) PeminjamanRepository {
	return &peminjamanRepository{db: db}
}

func (r *peminjamanRepository) GetAllPeminjaman() ([]entity.Peminjaman, error) {
	var peminjaman []entity.Peminjaman
	if err := r.db.Preload("Anggota", "id_anggota").Preload("Petugas", "id_petugas").Preload("DetailPeminjaman.Buku").Find(&peminjaman).Error; err != nil {
		return nil, err
	}
	return peminjaman, nil
}

func (r *peminjamanRepository) GetPeminjamanByID(id int) (entity.Peminjaman, error) {
	var peminjaman entity.Peminjaman
	if err := r.db.Preload("Anggota", "id_anggota").Preload("Petugas", "id_petugas").Preload("DetailPeminjaman.Buku").First(&peminjaman, "id_peminjaman = ?", id).Error; err != nil {
		return entity.Peminjaman{}, err
	}
	return peminjaman, nil
}

func (r *peminjamanRepository) CreatePeminjaman(p entity.Peminjaman) (entity.Peminjaman, error) {
	if err := r.db.Create(&p).Error; err != nil {
		return entity.Peminjaman{}, err
	}
	return r.GetPeminjamanByID(p.IDPeminjaman)
}

func (r *peminjamanRepository) UpdatePeminjaman(id int, p entity.Peminjaman) (entity.Peminjaman, error) {
	p.IDPeminjaman = id
	if err := r.db.Save(&p).Error; err != nil {
		return entity.Peminjaman{}, err
	}
	return r.GetPeminjamanByID(id)
}

func (r *peminjamanRepository) DeletePeminjaman(id int) error {
	return r.db.Delete(&entity.Peminjaman{}, "id_peminjaman = ?", id).Error
}
