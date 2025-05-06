package usecase

import (
    "perpustakaan/internal/entity"
    "perpustakaan/internal/repository"
)

type KategoriBukuUsecase interface {
    GetAllKategoriBuku() ([]entity.KategoriBuku, error)
    GetKategoriBukuByID(id int) (entity.KategoriBuku, error)
    CreateKategoriBuku(k entity.KategoriBuku) (entity.KategoriBuku, error)
    UpdateKategoriBuku(id int, k entity.KategoriBuku) (entity.KategoriBuku, error)
    DeleteKategoriBuku(id int) error
}

type kategoriBukuUsecase struct {
    repo repository.KategoriBukuRepository
}

func NewKategoriBukuUsecase(repo repository.KategoriBukuRepository) KategoriBukuUsecase {
    return &kategoriBukuUsecase{repo: repo}
}

func (u *kategoriBukuUsecase) GetAllKategoriBuku() ([]entity.KategoriBuku, error) {
    return u.repo.GetAllKategoriBuku()
}

func (u *kategoriBukuUsecase) GetKategoriBukuByID(id int) (entity.KategoriBuku, error) {
    return u.repo.GetKategoriBukuByID(id)
}

func (u *kategoriBukuUsecase) CreateKategoriBuku(k entity.KategoriBuku) (entity.KategoriBuku, error) {
    return u.repo.CreateKategoriBuku(k)
}

func (u *kategoriBukuUsecase) UpdateKategoriBuku(id int, k entity.KategoriBuku) (entity.KategoriBuku, error) {
    return u.repo.UpdateKategoriBuku(id, k)
}

func (u *kategoriBukuUsecase) DeleteKategoriBuku(id int) error {
    return u.repo.DeleteKategoriBuku(id)
} 