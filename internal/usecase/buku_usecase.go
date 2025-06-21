package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type BukuUsecase struct {
	repo *repository.BukuRepository
}

func NewBukuUsecase(repo *repository.BukuRepository) *BukuUsecase {
	return &BukuUsecase{repo: repo}
}

func (u *BukuUsecase) Create(ctx context.Context, buku *entity.Buku) error {
	return u.repo.Create(ctx, buku)
}

func (u *BukuUsecase) Update(ctx context.Context, buku *entity.Buku) error {
	return u.repo.Update(ctx, buku)
}

func (u *BukuUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *BukuUsecase) GetByID(ctx context.Context, id int64) (*entity.Buku, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *BukuUsecase) GetAll(ctx context.Context) ([]entity.Buku, error) {
	return u.repo.GetAll(ctx)
} 