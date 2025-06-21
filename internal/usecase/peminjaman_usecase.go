package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type PeminjamanUsecase struct {
	repo *repository.PeminjamanRepository
}

func NewPeminjamanUsecase(repo *repository.PeminjamanRepository) *PeminjamanUsecase {
	return &PeminjamanUsecase{repo: repo}
}

func (u *PeminjamanUsecase) Create(ctx context.Context, peminjaman *entity.Peminjaman) error {
	return u.repo.Create(ctx, peminjaman)
}

func (u *PeminjamanUsecase) Update(ctx context.Context, peminjaman *entity.Peminjaman) error {
	return u.repo.Update(ctx, peminjaman)
}

func (u *PeminjamanUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *PeminjamanUsecase) GetByID(ctx context.Context, id int64) (*entity.Peminjaman, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *PeminjamanUsecase) GetAll(ctx context.Context) ([]entity.Peminjaman, error) {
	return u.repo.GetAll(ctx)
} 