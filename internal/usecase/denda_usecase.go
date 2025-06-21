package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type DendaUsecase struct {
	repo *repository.DendaRepository
}

func NewDendaUsecase(repo *repository.DendaRepository) *DendaUsecase {
	return &DendaUsecase{repo: repo}
}

func (u *DendaUsecase) Create(ctx context.Context, denda *entity.Denda) error {
	return u.repo.Create(ctx, denda)
}

func (u *DendaUsecase) Update(ctx context.Context, denda *entity.Denda) error {
	return u.repo.Update(ctx, denda)
}

func (u *DendaUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *DendaUsecase) GetByID(ctx context.Context, id int64) (*entity.Denda, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *DendaUsecase) GetAll(ctx context.Context) ([]entity.Denda, error) {
	return u.repo.GetAll(ctx)
} 