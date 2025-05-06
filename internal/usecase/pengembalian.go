package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type PengembalianUsecase interface {
    GetAllPengembalian() ([]entity.Pengembalian, error)
    GetPengembalianByID(id int) (entity.Pengembalian, error)
    CreatePengembalian(p entity.Pengembalian) (entity.Pengembalian, error)
    UpdatePengembalian(id int, p entity.Pengembalian) (entity.Pengembalian, error)
    DeletePengembalian(id int) error
}

type pengembalianUsecase struct {
    repo repository.PengembalianRepository
}

func NewPengembalianUsecase(repo repository.PengembalianRepository) PengembalianUsecase {
    return &pengembalianUsecase{repo: repo}
}

func (u *pengembalianUsecase) GetAllPengembalian() ([]entity.Pengembalian, error) {
    return u.repo.GetAllPengembalian()
}

func (u *pengembalianUsecase) GetPengembalianByID(id int) (entity.Pengembalian, error) {
    return u.repo.GetPengembalianByID(id)
}

func (u *pengembalianUsecase) CreatePengembalian(p entity.Pengembalian) (entity.Pengembalian, error) {
    return u.repo.CreatePengembalian(p)
}

func (u *pengembalianUsecase) UpdatePengembalian(id int, p entity.Pengembalian) (entity.Pengembalian, error) {
    return u.repo.UpdatePengembalian(id, p)
}

func (u *pengembalianUsecase) DeletePengembalian(id int) error {
    return u.repo.DeletePengembalian(id)
} 