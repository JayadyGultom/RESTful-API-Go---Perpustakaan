package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type BukuUsecase interface {
    GetAllBuku() ([]entity.Buku, error)
    GetBukuByID(id int) (entity.Buku, error)
    CreateBuku(b entity.Buku) (entity.Buku, error)
    UpdateBuku(id int, b entity.Buku) (entity.Buku, error)
    DeleteBuku(id int) error
}

type bukuUsecase struct {
    repo repository.BukuRepository
}

func NewBukuUsecase(repo repository.BukuRepository) BukuUsecase {
    return &bukuUsecase{repo: repo}
}

func (u *bukuUsecase) GetAllBuku() ([]entity.Buku, error) {
    return u.repo.GetAllBuku()
}

func (u *bukuUsecase) GetBukuByID(id int) (entity.Buku, error) {
    return u.repo.GetBukuByID(id)
}

func (u *bukuUsecase) CreateBuku(b entity.Buku) (entity.Buku, error) {
    return u.repo.CreateBuku(b)
}

func (u *bukuUsecase) UpdateBuku(id int, b entity.Buku) (entity.Buku, error) {
    return u.repo.UpdateBuku(id, b)
}

func (u *bukuUsecase) DeleteBuku(id int) error {
    return u.repo.DeleteBuku(id)
} 