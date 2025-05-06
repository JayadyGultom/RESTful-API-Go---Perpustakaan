package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type PenulisUsecase interface {
    GetAllPenulis() ([]entity.Penulis, error)
    GetPenulisByID(id int) (entity.Penulis, error)
    CreatePenulis(p entity.Penulis) (entity.Penulis, error)
    UpdatePenulis(id int, p entity.Penulis) (entity.Penulis, error)
    DeletePenulis(id int) error
}

type penulisUsecase struct {
    repo repository.PenulisRepository
}

func NewPenulisUsecase(repo repository.PenulisRepository) PenulisUsecase {
    return &penulisUsecase{repo: repo}
}

func (u *penulisUsecase) GetAllPenulis() ([]entity.Penulis, error) {
    return u.repo.GetAllPenulis()
}

func (u *penulisUsecase) GetPenulisByID(id int) (entity.Penulis, error) {
    return u.repo.GetPenulisByID(id)
}

func (u *penulisUsecase) CreatePenulis(p entity.Penulis) (entity.Penulis, error) {
    return u.repo.CreatePenulis(p)
}

func (u *penulisUsecase) UpdatePenulis(id int, p entity.Penulis) (entity.Penulis, error) {
    return u.repo.UpdatePenulis(id, p)
}

func (u *penulisUsecase) DeletePenulis(id int) error {
    return u.repo.DeletePenulis(id)
} 