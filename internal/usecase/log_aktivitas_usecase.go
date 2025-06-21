package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type LogAktivitasUsecase struct {
	repo *repository.LogAktivitasRepository
}

func NewLogAktivitasUsecase(repo *repository.LogAktivitasRepository) *LogAktivitasUsecase {
	return &LogAktivitasUsecase{repo: repo}
}

func (u *LogAktivitasUsecase) Create(ctx context.Context, log *entity.LogAktivitas) error {
	return u.repo.Create(ctx, log)
}

func (u *LogAktivitasUsecase) Update(ctx context.Context, log *entity.LogAktivitas) error {
	return u.repo.Update(ctx, log)
}

func (u *LogAktivitasUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *LogAktivitasUsecase) GetByID(ctx context.Context, id int64) (*entity.LogAktivitas, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *LogAktivitasUsecase) GetAll(ctx context.Context) ([]entity.LogAktivitas, error) {
	return u.repo.GetAll(ctx)
} 