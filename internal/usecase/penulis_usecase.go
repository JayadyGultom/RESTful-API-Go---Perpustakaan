package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type PenulisUsecase struct {
	repo *repository.PenulisRepository
}

func NewPenulisUsecase(repo *repository.PenulisRepository) *PenulisUsecase {
	return &PenulisUsecase{repo: repo}
}

func (u *PenulisUsecase) Create(ctx context.Context, penulis *entity.Penulis) error {
	return u.repo.Create(ctx, penulis)
}

func (u *PenulisUsecase) Update(ctx context.Context, penulis *entity.Penulis) error {
	return u.repo.Update(ctx, penulis)
}

func (u *PenulisUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *PenulisUsecase) GetByID(ctx context.Context, id int64) (*entity.Penulis, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *PenulisUsecase) GetAll(ctx context.Context) ([]entity.Penulis, error) {
	return u.repo.GetAll(ctx)
}

func (u *PenulisUsecase) GetByNama(ctx context.Context, nama string) (*entity.Penulis, error) {
	return u.repo.GetByNama(ctx, nama)
} 