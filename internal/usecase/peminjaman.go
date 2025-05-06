package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type PeminjamanUsecase interface {
    GetAllPeminjaman() ([]entity.Peminjaman, error)
    GetPeminjamanByID(id int) (entity.Peminjaman, error)
    CreatePeminjaman(p entity.Peminjaman) (entity.Peminjaman, error)
    UpdatePeminjaman(id int, p entity.Peminjaman) (entity.Peminjaman, error)
    DeletePeminjaman(id int) error
}

type peminjamanUsecase struct {
    repo repository.PeminjamanRepository
}

func NewPeminjamanUsecase(repo repository.PeminjamanRepository) PeminjamanUsecase {
    return &peminjamanUsecase{repo: repo}
}

func (u *peminjamanUsecase) GetAllPeminjaman() ([]entity.Peminjaman, error) {
    return u.repo.GetAllPeminjaman()
}

func (u *peminjamanUsecase) GetPeminjamanByID(id int) (entity.Peminjaman, error) {
    return u.repo.GetPeminjamanByID(id)
}

func (u *peminjamanUsecase) CreatePeminjaman(p entity.Peminjaman) (entity.Peminjaman, error) {
    return u.repo.CreatePeminjaman(p)
}

func (u *peminjamanUsecase) UpdatePeminjaman(id int, p entity.Peminjaman) (entity.Peminjaman, error) {
    return u.repo.UpdatePeminjaman(id, p)
}

func (u *peminjamanUsecase) DeletePeminjaman(id int) error {
    return u.repo.DeletePeminjaman(id)
} 