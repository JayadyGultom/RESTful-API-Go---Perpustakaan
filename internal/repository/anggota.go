package repository

import (
    "perpustakaan/internal/entity"
    "gorm.io/gorm"
)

type AnggotaRepository interface {
    GetAllAnggota() ([]entity.Anggota, error)
    GetAnggotaByID(id int) (entity.Anggota, error)
    CreateAnggota(a entity.Anggota) (entity.Anggota, error)
    UpdateAnggota(id int, a entity.Anggota) (entity.Anggota, error)
    DeleteAnggota(id int) error
}

type anggotaRepository struct {
    db *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) AnggotaRepository {
    return &anggotaRepository{db: db}
}

func (r *anggotaRepository) GetAllAnggota() ([]entity.Anggota, error) {
    var anggota []entity.Anggota
    if err := r.db.Find(&anggota).Error; err != nil {
        return nil, err
    }
    return anggota, nil
}

func (r *anggotaRepository) GetAnggotaByID(id int) (entity.Anggota, error) {
    var anggota entity.Anggota
    if err := r.db.First(&anggota, "id_anggota = ?", id).Error; err != nil {
        return entity.Anggota{}, err
    }
    return anggota, nil
}

func (r *anggotaRepository) CreateAnggota(a entity.Anggota) (entity.Anggota, error) {
    if err := r.db.Create(&a).Error; err != nil {
        return entity.Anggota{}, err
    }
    return a, nil
}

func (r *anggotaRepository) UpdateAnggota(id int, a entity.Anggota) (entity.Anggota, error) {
    a.IDAnggota = id
    if err := r.db.Save(&a).Error; err != nil {
        return entity.Anggota{}, err
    }
    return a, nil
}

func (r *anggotaRepository) DeleteAnggota(id int) error {
    return r.db.Delete(&entity.Anggota{}, "id_anggota = ?", id).Error
}
