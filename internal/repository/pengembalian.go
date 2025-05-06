package repository

import (
    "perpustakaan/internal/entity"
    "gorm.io/gorm"
)

type PengembalianRepository interface {
    GetAllPengembalian() ([]entity.Pengembalian, error)
    GetPengembalianByID(id int) (entity.Pengembalian, error)
    CreatePengembalian(p entity.Pengembalian) (entity.Pengembalian, error)
    UpdatePengembalian(id int, p entity.Pengembalian) (entity.Pengembalian, error)
    DeletePengembalian(id int) error
}

type pengembalianRepository struct {
    db *gorm.DB
}

func NewPengembalianRepository(db *gorm.DB) PengembalianRepository {
    return &pengembalianRepository{db: db}
}

func (r *pengembalianRepository) GetAllPengembalian() ([]entity.Pengembalian, error) {
    var pengembalian []entity.Pengembalian
    if err := r.db.Find(&pengembalian).Error; err != nil {
        return nil, err
    }
    return pengembalian, nil
}

func (r *pengembalianRepository) GetPengembalianByID(id int) (entity.Pengembalian, error) {
    var pengembalian entity.Pengembalian
    if err := r.db.First(&pengembalian, "id_pengembalian = ?", id).Error; err != nil {
        return entity.Pengembalian{}, err
    }
    return pengembalian, nil
}

func (r *pengembalianRepository) CreatePengembalian(p entity.Pengembalian) (entity.Pengembalian, error) {
    if err := r.db.Create(&p).Error; err != nil {
        return entity.Pengembalian{}, err
    }
    return r.GetPengembalianByID(p.IDPengembalian)
}

func (r *pengembalianRepository) UpdatePengembalian(id int, p entity.Pengembalian) (entity.Pengembalian, error) {
    p.IDPengembalian = id
    if err := r.db.Save(&p).Error; err != nil {
        return entity.Pengembalian{}, err
    }
    return r.GetPengembalianByID(id)
}

func (r *pengembalianRepository) DeletePengembalian(id int) error {
    return r.db.Delete(&entity.Pengembalian{}, "id_pengembalian = ?", id).Error
} 