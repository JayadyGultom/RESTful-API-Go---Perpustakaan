package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type PenerbitUsecase struct {
	repo *repository.PenerbitRepository
}

func NewPenerbitUsecase(repo *repository.PenerbitRepository) *PenerbitUsecase {
	return &PenerbitUsecase{repo: repo}
}

func (u *PenerbitUsecase) Create(ctx context.Context, penerbit *entity.Penerbit) error {
	return u.repo.Create(ctx, penerbit)
}

func (u *PenerbitUsecase) Update(ctx context.Context, penerbit *entity.Penerbit) error {
	return u.repo.Update(ctx, penerbit)
}

func (u *PenerbitUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *PenerbitUsecase) GetByID(ctx context.Context, id int64) (*entity.Penerbit, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *PenerbitUsecase) GetAll(ctx context.Context) ([]entity.Penerbit, error) {
	return u.repo.GetAll(ctx)
}

func (u *PenerbitUsecase) GetByNama(ctx context.Context, nama string) (*entity.Penerbit, error) {
	return u.repo.GetByNama(ctx, nama)
} 