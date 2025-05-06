package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type AnggotaUsecase interface {
    GetAllAnggota() ([]entity.Anggota, error)
    GetAnggotaByID(id int) (entity.Anggota, error)
    CreateAnggota(a entity.Anggota) (entity.Anggota, error)
    UpdateAnggota(id int, a entity.Anggota) (entity.Anggota, error)
    DeleteAnggota(id int) error
}

type anggotaUsecase struct {
    repo repository.AnggotaRepository
}

func NewAnggotaUsecase(repo repository.AnggotaRepository) AnggotaUsecase {
    return &anggotaUsecase{repo: repo}
}

func (u *anggotaUsecase) GetAllAnggota() ([]entity.Anggota, error) {
    return u.repo.GetAllAnggota()
}

func (u *anggotaUsecase) GetAnggotaByID(id int) (entity.Anggota, error) {
    return u.repo.GetAnggotaByID(id)
}

func (u *anggotaUsecase) CreateAnggota(a entity.Anggota) (entity.Anggota, error) {
    return u.repo.CreateAnggota(a)
}

func (u *anggotaUsecase) UpdateAnggota(id int, a entity.Anggota) (entity.Anggota, error) {
    return u.repo.UpdateAnggota(id, a)
}

func (u *anggotaUsecase) DeleteAnggota(id int) error {
    return u.repo.DeleteAnggota(id)
}
