package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type KategoriBukuUsecase struct {
	repo *repository.KategoriBukuRepository
}

func NewKategoriBukuUsecase(repo *repository.KategoriBukuRepository) *KategoriBukuUsecase {
	return &KategoriBukuUsecase{repo: repo}
}

func (u *KategoriBukuUsecase) Create(ctx context.Context, kategori *entity.KategoriBuku) error {
	return u.repo.Create(ctx, kategori)
}

func (u *KategoriBukuUsecase) Update(ctx context.Context, kategori *entity.KategoriBuku) error {
	return u.repo.Update(ctx, kategori)
}

func (u *KategoriBukuUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *KategoriBukuUsecase) GetByID(ctx context.Context, id int64) (*entity.KategoriBuku, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *KategoriBukuUsecase) GetAll(ctx context.Context) ([]entity.KategoriBuku, error) {
	return u.repo.GetAll(ctx)
}

func (u *KategoriBukuUsecase) GetByNama(ctx context.Context, nama string) (*entity.KategoriBuku, error) {
	return u.repo.GetByNama(ctx, nama)
} 