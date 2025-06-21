package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type KategoriBukuController struct {
	kategoriUsecase *usecase.KategoriBukuUsecase
}

func NewKategoriBukuController(kategoriUsecase *usecase.KategoriBukuUsecase) *KategoriBukuController {
	return &KategoriBukuController{kategoriUsecase: kategoriUsecase}
}

func (c *KategoriBukuController) Create(w http.ResponseWriter, r *http.Request) {
	var kategori entity.KategoriBuku
	if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.kategoriUsecase.Create(r.Context(), &kategori); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategori)
}

func (c *KategoriBukuController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var kategori entity.KategoriBuku
	if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kategori.IDKategori = id
	if err := c.kategoriUsecase.Update(r.Context(), &kategori); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func (c *KategoriBukuController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.kategoriUsecase.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori berhasil dihapus"})
}

func (c *KategoriBukuController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	kategori, err := c.kategoriUsecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func (c *KategoriBukuController) GetAll(w http.ResponseWriter, r *http.Request) {
	kategoris, err := c.kategoriUsecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategoris)
}

func (c *KategoriBukuController) GetByNama(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("nama")
	if nama == "" {
		http.Error(w, "Nama kategori tidak boleh kosong", http.StatusBadRequest)
		return
	}

	kategori, err := c.kategoriUsecase.GetByNama(r.Context(), nama)
	if err != nil {
		http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
} 