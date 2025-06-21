package usecase

import (
	"context"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type PetugasUsecase struct {
	repo *repository.PetugasRepository
}

func NewPetugasUsecase(repo *repository.PetugasRepository) *PetugasUsecase {
	return &PetugasUsecase{repo: repo}
}

func (u *PetugasUsecase) Create(ctx context.Context, petugas *entity.Petugas) error {
	return u.repo.Create(ctx, petugas)
}

func (u *PetugasUsecase) Update(ctx context.Context, petugas *entity.Petugas) error {
	return u.repo.Update(ctx, petugas)
}

func (u *PetugasUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *PetugasUsecase) GetByID(ctx context.Context, id int64) (*entity.Petugas, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *PetugasUsecase) GetAll(ctx context.Context) ([]entity.Petugas, error) {
	return u.repo.GetAll(ctx)
}

func (u *PetugasUsecase) GetByUsername(ctx context.Context, username string) (*entity.Petugas, error) {
	return u.repo.GetByUsername(ctx, username)
} 