package repository

import (
	"fmt"
	"perpustakaan/internal/entity"
	"time"
	"gorm.io/gorm"
)

type AnggotaRepository interface {
	GetAll() ([]entity.Anggota, error)
	GetByID(id int) (entity.Anggota, error)
	Create(a entity.Anggota) error
	Update(a entity.Anggota) error
	Delete(id int) error
	GenerateNomorIdentitas() (string, error)
}

type anggotaRepository struct {
	db *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) AnggotaRepository {
	return &anggotaRepository{db}
}

func (r *anggotaRepository) GetAll() ([]entity.Anggota, error) {
	var anggotas []entity.Anggota
	result := r.db.Find(&anggotas)
	return anggotas, result.Error
}

func (r *anggotaRepository) GetByID(id int) (entity.Anggota, error) {
	var anggota entity.Anggota
	result := r.db.First(&anggota, id)
	return anggota, result.Error
}

func (r *anggotaRepository) Create(a entity.Anggota) error {
	// Generate nomor identitas jika kosong
	if a.NomorIdentitas == "" {
		nomorIdentitas, err := r.GenerateNomorIdentitas()
		if err != nil {
			return err
		}
		a.NomorIdentitas = nomorIdentitas
	}
	
	// Set tanggal bergabung jika kosong
	if a.TanggalBergabung.IsZero() {
		a.TanggalBergabung = time.Now()
	}
	
	result := r.db.Create(&a)
	return result.Error
}

func (r *anggotaRepository) Update(a entity.Anggota) error {
	result := r.db.Save(&a)
	return result.Error
}

func (r *anggotaRepository) Delete(id int) error {
	result := r.db.Delete(&entity.Anggota{}, id)
	return result.Error
}

// GenerateNomorIdentitas generates a unique nomor identitas
func (r *anggotaRepository) GenerateNomorIdentitas() (string, error) {
	// Hitung total anggota yang sudah ada
	var count int64
	err := r.db.Model(&entity.Anggota{}).Count(&count).Error
	if err != nil {
		return "", err
	}
	
	// Generate nomor urut dengan padding 4 digit
	sequence := fmt.Sprintf("%04d", count+1)
	nomorIdentitas := fmt.Sprintf("JAY-%s", sequence)
	
	return nomorIdentitas, nil
}