package repository

import (
    "perpustakaan/internal/entity"
    "gorm.io/gorm"
)

type PenerbitRepository interface {
    GetAllPenerbit() ([]entity.Penerbit, error)
    GetPenerbitByID(id int) (entity.Penerbit, error)
    CreatePenerbit(p entity.Penerbit) (entity.Penerbit, error)
    UpdatePenerbit(id int, p entity.Penerbit) (entity.Penerbit, error)
    DeletePenerbit(id int) error
}

type penerbitRepository struct {
    db *gorm.DB
}

func NewPenerbitRepository(db *gorm.DB) PenerbitRepository {
    return &penerbitRepository{db: db}
}

func (r *penerbitRepository) GetAllPenerbit() ([]entity.Penerbit, error) {
    var penerbit []entity.Penerbit
    if err := r.db.Find(&penerbit).Error; err != nil {
        return nil, err
    }
    return penerbit, nil
}

func (r *penerbitRepository) GetPenerbitByID(id int) (entity.Penerbit, error) {
    var penerbit entity.Penerbit
    if err := r.db.First(&penerbit, "id_penerbit = ?", id).Error; err != nil {
        return entity.Penerbit{}, err
    }
    return penerbit, nil
}

func (r *penerbitRepository) CreatePenerbit(p entity.Penerbit) (entity.Penerbit, error) {
    if err := r.db.Create(&p).Error; err != nil {
        return entity.Penerbit{}, err
    }
    return p, nil
}

func (r *penerbitRepository) UpdatePenerbit(id int, p entity.Penerbit) (entity.Penerbit, error) {
    p.IDPenerbit = id
    if err := r.db.Save(&p).Error; err != nil {
        return entity.Penerbit{}, err
    }
    return p, nil
}

func (r *penerbitRepository) DeletePenerbit(id int) error {
    return r.db.Delete(&entity.Penerbit{}, "id_penerbit = ?", id).Error
} 