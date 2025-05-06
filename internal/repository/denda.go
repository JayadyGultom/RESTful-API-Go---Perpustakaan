package repository

import (
    "perpustakaan/internal/entity"
    "gorm.io/gorm"
)

type DendaRepository interface {
    GetAllDenda() ([]entity.Denda, error)
    GetDendaByID(id int) (entity.Denda, error)
    CreateDenda(d entity.Denda) (entity.Denda, error)
    UpdateDenda(id int, d entity.Denda) (entity.Denda, error)
    DeleteDenda(id int) error
}

type dendaRepository struct {
    db *gorm.DB
}

func NewDendaRepository(db *gorm.DB) DendaRepository {
    return &dendaRepository{db: db}
}

func (r *dendaRepository) GetAllDenda() ([]entity.Denda, error) {
    var denda []entity.Denda
    if err := r.db.Preload("Pengembalian").Preload("Pengembalian.Peminjaman").Find(&denda).Error; err != nil {
        return nil, err
    }
    return denda, nil
}

func (r *dendaRepository) GetDendaByID(id int) (entity.Denda, error) {
    var denda entity.Denda
    if err := r.db.Preload("Pengembalian").First(&denda, "id_denda = ?", id).Error; err != nil {
        return entity.Denda{}, err
    }
    return denda, nil
}

func (r *dendaRepository) CreateDenda(d entity.Denda) (entity.Denda, error) {
    if err := r.db.Create(&d).Error; err != nil {
        return entity.Denda{}, err
    }
    return r.GetDendaByID(d.IDDenda)
}

func (r *dendaRepository) UpdateDenda(id int, d entity.Denda) (entity.Denda, error) {
    d.IDDenda = id
    if err := r.db.Save(&d).Error; err != nil {
        return entity.Denda{}, err
    }
    return r.GetDendaByID(id)
}

func (r *dendaRepository) DeleteDenda(id int) error {
    return r.db.Delete(&entity.Denda{}, "id_denda = ?", id).Error
} 