package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type DetailPeminjamanUsecase interface {
    GetAllDetailPeminjaman() ([]entity.DetailPeminjaman, error)
    GetDetailPeminjamanByID(id int) (entity.DetailPeminjaman, error)
    CreateDetailPeminjaman(d entity.DetailPeminjaman) (entity.DetailPeminjaman, error)
    UpdateDetailPeminjaman(id int, d entity.DetailPeminjaman) (entity.DetailPeminjaman, error)
    DeleteDetailPeminjaman(id int) error
}

type detailPeminjamanUsecase struct {
    repo repository.DetailPeminjamanRepository
}

func NewDetailPeminjamanUsecase(repo repository.DetailPeminjamanRepository) DetailPeminjamanUsecase {
    return &detailPeminjamanUsecase{repo: repo}
}

func (u *detailPeminjamanUsecase) GetAllDetailPeminjaman() ([]entity.DetailPeminjaman, error) {
    return u.repo.GetAllDetailPeminjaman()
}

func (u *detailPeminjamanUsecase) GetDetailPeminjamanByID(id int) (entity.DetailPeminjaman, error) {
    return u.repo.GetDetailPeminjamanByID(id)
}

func (u *detailPeminjamanUsecase) CreateDetailPeminjaman(d entity.DetailPeminjaman) (entity.DetailPeminjaman, error) {
    return u.repo.CreateDetailPeminjaman(d)
}

func (u *detailPeminjamanUsecase) UpdateDetailPeminjaman(id int, d entity.DetailPeminjaman) (entity.DetailPeminjaman, error) {
    return u.repo.UpdateDetailPeminjaman(id, d)
}

func (u *detailPeminjamanUsecase) DeleteDetailPeminjaman(id int) error {
    return u.repo.DeleteDetailPeminjaman(id)
} 