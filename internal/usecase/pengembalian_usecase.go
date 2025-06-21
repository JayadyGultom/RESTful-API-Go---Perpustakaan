package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type PengembalianUsecase struct {
	repo *repository.PengembalianRepository
}

func NewPengembalianUsecase(repo *repository.PengembalianRepository) *PengembalianUsecase {
	return &PengembalianUsecase{repo: repo}
}

func (u *PengembalianUsecase) Create(ctx context.Context, pengembalian *entity.Pengembalian) error {
	return u.repo.Create(ctx, pengembalian)
}

func (u *PengembalianUsecase) Update(ctx context.Context, pengembalian *entity.Pengembalian) error {
	return u.repo.Update(ctx, pengembalian)
}

func (u *PengembalianUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *PengembalianUsecase) GetByID(ctx context.Context, id int64) (*entity.Pengembalian, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *PengembalianUsecase) GetAll(ctx context.Context) ([]entity.Pengembalian, error) {
	return u.repo.GetAll(ctx)
} 