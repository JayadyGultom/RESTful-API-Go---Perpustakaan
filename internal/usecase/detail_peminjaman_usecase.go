package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type DetailPeminjamanUsecase struct {
	repo *repository.DetailPeminjamanRepository
}

func NewDetailPeminjamanUsecase(repo *repository.DetailPeminjamanRepository) *DetailPeminjamanUsecase {
	return &DetailPeminjamanUsecase{repo: repo}
}

func (u *DetailPeminjamanUsecase) Create(ctx context.Context, detail *entity.DetailPeminjaman) error {
	return u.repo.Create(ctx, detail)
}

func (u *DetailPeminjamanUsecase) Update(ctx context.Context, detail *entity.DetailPeminjaman) error {
	return u.repo.Update(ctx, detail)
}

func (u *DetailPeminjamanUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *DetailPeminjamanUsecase) GetByID(ctx context.Context, id int64) (*entity.DetailPeminjaman, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *DetailPeminjamanUsecase) GetAll(ctx context.Context) ([]entity.DetailPeminjaman, error) {
	return u.repo.GetAll(ctx)
} 