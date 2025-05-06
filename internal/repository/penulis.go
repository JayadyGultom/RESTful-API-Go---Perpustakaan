package repository

import (
    "perpustakaan/internal/entity"
    "gorm.io/gorm"
)

type PenulisRepository interface {
    GetAllPenulis() ([]entity.Penulis, error)
    GetPenulisByID(id int) (entity.Penulis, error)
    CreatePenulis(p entity.Penulis) (entity.Penulis, error)
    UpdatePenulis(id int, p entity.Penulis) (entity.Penulis, error)
    DeletePenulis(id int) error
}

type penulisRepository struct {
    db *gorm.DB
}

func NewPenulisRepository(db *gorm.DB) PenulisRepository {
    return &penulisRepository{db: db}
}

func (r *penulisRepository) GetAllPenulis() ([]entity.Penulis, error) {
    var penulis []entity.Penulis
    if err := r.db.Find(&penulis).Error; err != nil {
        return nil, err
    }
    return penulis, nil
}

func (r *penulisRepository) GetPenulisByID(id int) (entity.Penulis, error) {
    var penulis entity.Penulis
    if err := r.db.First(&penulis, "id_penulis = ?", id).Error; err != nil {
        return entity.Penulis{}, err
    }
    return penulis, nil
}

func (r *penulisRepository) CreatePenulis(p entity.Penulis) (entity.Penulis, error) {
    if err := r.db.Create(&p).Error; err != nil {
        return entity.Penulis{}, err
    }
    return p, nil
}

func (r *penulisRepository) UpdatePenulis(id int, p entity.Penulis) (entity.Penulis, error) {
    p.IDPenulis = id
    if err := r.db.Save(&p).Error; err != nil {
        return entity.Penulis{}, err
    }
    return p, nil
}

func (r *penulisRepository) DeletePenulis(id int) error {
    return r.db.Delete(&entity.Penulis{}, "id_penulis = ?", id).Error
} 