package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type ReservasiUsecase struct {
	repo *repository.ReservasiRepository
}

func NewReservasiUsecase(repo *repository.ReservasiRepository) *ReservasiUsecase {
	return &ReservasiUsecase{repo: repo}
}

func (u *ReservasiUsecase) Create(ctx context.Context, reservasi *entity.Reservasi) error {
	return u.repo.Create(ctx, reservasi)
}

func (u *ReservasiUsecase) Update(ctx context.Context, reservasi *entity.Reservasi) error {
	return u.repo.Update(ctx, reservasi)
}

func (u *ReservasiUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *ReservasiUsecase) GetByID(ctx context.Context, id int64) (*entity.Reservasi, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *ReservasiUsecase) GetAll(ctx context.Context) ([]entity.Reservasi, error) {
	return u.repo.GetAll(ctx)
} 