package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type DendaUsecase interface {
    GetAllDenda() ([]entity.Denda, error)
    GetDendaByID(id int) (entity.Denda, error)
    CreateDenda(d entity.Denda) (entity.Denda, error)
    UpdateDenda(id int, d entity.Denda) (entity.Denda, error)
    DeleteDenda(id int) error
}

type dendaUsecase struct {
    repo repository.DendaRepository
}

func NewDendaUsecase(repo repository.DendaRepository) DendaUsecase {
    return &dendaUsecase{repo: repo}
}

func (u *dendaUsecase) GetAllDenda() ([]entity.Denda, error) {
    return u.repo.GetAllDenda()
}

func (u *dendaUsecase) GetDendaByID(id int) (entity.Denda, error) {
    return u.repo.GetDendaByID(id)
}

func (u *dendaUsecase) CreateDenda(d entity.Denda) (entity.Denda, error) {
    return u.repo.CreateDenda(d)
}

func (u *dendaUsecase) UpdateDenda(id int, d entity.Denda) (entity.Denda, error) {
    return u.repo.UpdateDenda(id, d)
}

func (u *dendaUsecase) DeleteDenda(id int) error {
    return u.repo.DeleteDenda(id)
} 