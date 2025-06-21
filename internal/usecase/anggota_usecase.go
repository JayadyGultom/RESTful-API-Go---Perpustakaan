package usecase

import (
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type AnggotaUsecase interface {
	GetAll() ([]entity.Anggota, error)
	GetByID(id int) (entity.Anggota, error)
	Create(a entity.Anggota) error
	Update(a entity.Anggota) error
	Delete(id int) error
}

type anggotaUsecase struct {
	repo repository.AnggotaRepository
}

func NewAnggotaUsecase(r repository.AnggotaRepository) AnggotaUsecase {
	return &anggotaUsecase{repo: r}
}

func (u *anggotaUsecase) GetAll() ([]entity.Anggota, error) {
	return u.repo.GetAll()
}

func (u *anggotaUsecase) GetByID(id int) (entity.Anggota, error) {
	return u.repo.GetByID(id)
}

func (u *anggotaUsecase) Create(a entity.Anggota) error {
	return u.repo.Create(a)
}

func (u *anggotaUsecase) Update(a entity.Anggota) error {
	return u.repo.Update(a)
}

func (u *anggotaUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}