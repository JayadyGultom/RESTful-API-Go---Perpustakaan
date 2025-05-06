package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type PetugasUsecase interface {
    GetAllPetugas() ([]entity.Petugas, error)
    GetPetugasByID(id int) (entity.Petugas, error)
    CreatePetugas(p entity.Petugas) (entity.Petugas, error)
    UpdatePetugas(id int, p entity.Petugas) (entity.Petugas, error)
    DeletePetugas(id int) error
}

type petugasUsecase struct {
    repo repository.PetugasRepository
}

func NewPetugasUsecase(repo repository.PetugasRepository) PetugasUsecase {
    return &petugasUsecase{repo: repo}
}

func (u *petugasUsecase) GetAllPetugas() ([]entity.Petugas, error) {
    return u.repo.GetAllPetugas()
}

func (u *petugasUsecase) GetPetugasByID(id int) (entity.Petugas, error) {
    return u.repo.GetPetugasByID(id)
}

func (u *petugasUsecase) CreatePetugas(p entity.Petugas) (entity.Petugas, error) {
    return u.repo.CreatePetugas(p)
}

func (u *petugasUsecase) UpdatePetugas(id int, p entity.Petugas) (entity.Petugas, error) {
    return u.repo.UpdatePetugas(id, p)
}

func (u *petugasUsecase) DeletePetugas(id int) error {
    return u.repo.DeletePetugas(id)
} 