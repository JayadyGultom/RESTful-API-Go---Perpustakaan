package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type PenerbitUsecase interface {
    GetAllPenerbit() ([]entity.Penerbit, error)
    GetPenerbitByID(id int) (entity.Penerbit, error)
    CreatePenerbit(p entity.Penerbit) (entity.Penerbit, error)
    UpdatePenerbit(id int, p entity.Penerbit) (entity.Penerbit, error)
    DeletePenerbit(id int) error
}

type penerbitUsecase struct {
    repo repository.PenerbitRepository
}

func NewPenerbitUsecase(repo repository.PenerbitRepository) PenerbitUsecase {
    return &penerbitUsecase{repo: repo}
}

func (u *penerbitUsecase) GetAllPenerbit() ([]entity.Penerbit, error) {
    return u.repo.GetAllPenerbit()
}

func (u *penerbitUsecase) GetPenerbitByID(id int) (entity.Penerbit, error) {
    return u.repo.GetPenerbitByID(id)
}

func (u *penerbitUsecase) CreatePenerbit(p entity.Penerbit) (entity.Penerbit, error) {
    return u.repo.CreatePenerbit(p)
}

func (u *penerbitUsecase) UpdatePenerbit(id int, p entity.Penerbit) (entity.Penerbit, error) {
    return u.repo.UpdatePenerbit(id, p)
}

func (u *penerbitUsecase) DeletePenerbit(id int) error {
    return u.repo.DeletePenerbit(id)
} 