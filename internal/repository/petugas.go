package repository

import (
    "perpustakaan/internal/entity"
    "gorm.io/gorm"
)

type PetugasRepository interface {
    GetAllPetugas() ([]entity.Petugas, error)
    GetPetugasByID(id int) (entity.Petugas, error)
    CreatePetugas(p entity.Petugas) (entity.Petugas, error)
    UpdatePetugas(id int, p entity.Petugas) (entity.Petugas, error)
    DeletePetugas(id int) error
}

type petugasRepository struct {
    db *gorm.DB
}

func NewPetugasRepository(db *gorm.DB) PetugasRepository {
    return &petugasRepository{db: db}
}

func (r *petugasRepository) GetAllPetugas() ([]entity.Petugas, error) {
    var petugas []entity.Petugas
    if err := r.db.Find(&petugas).Error; err != nil {
        return nil, err
    }
    return petugas, nil
}

func (r *petugasRepository) GetPetugasByID(id int) (entity.Petugas, error) {
    var petugas entity.Petugas
    if err := r.db.First(&petugas, "id_petugas = ?", id).Error; err != nil {
        return entity.Petugas{}, err
    }
    return petugas, nil
}

func (r *petugasRepository) CreatePetugas(p entity.Petugas) (entity.Petugas, error) {
    if err := r.db.Create(&p).Error; err != nil {
        return entity.Petugas{}, err
    }
    return p, nil
}

func (r *petugasRepository) UpdatePetugas(id int, p entity.Petugas) (entity.Petugas, error) {
    p.IDPetugas = id
    if err := r.db.Save(&p).Error; err != nil {
        return entity.Petugas{}, err
    }
    return p, nil
}

func (r *petugasRepository) DeletePetugas(id int) error {
    return r.db.Delete(&entity.Petugas{}, "id_petugas = ?", id).Error
} 