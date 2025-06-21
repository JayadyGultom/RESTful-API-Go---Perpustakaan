package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type PeminjamanController struct {
	usecase *usecase.PeminjamanUsecase
}

func NewPeminjamanController(usecase *usecase.PeminjamanUsecase) *PeminjamanController {
	return &PeminjamanController{usecase: usecase}
}

func (c *PeminjamanController) Create(w http.ResponseWriter, r *http.Request) {
	var peminjaman entity.Peminjaman
	if err := json.NewDecoder(r.Body).Decode(&peminjaman); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.usecase.Create(r.Context(), &peminjaman); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(peminjaman)
}

func (c *PeminjamanController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var peminjaman entity.Peminjaman
	if err := json.NewDecoder(r.Body).Decode(&peminjaman); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	peminjaman.IDPeminjaman = id
	if err := c.usecase.Update(r.Context(), &peminjaman); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peminjaman)
}

func (c *PeminjamanController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := c.usecase.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Peminjaman berhasil dihapus"})
}

func (c *PeminjamanController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	peminjaman, err := c.usecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Peminjaman tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peminjaman)
}

func (c *PeminjamanController) GetAll(w http.ResponseWriter, r *http.Request) {
	peminjamans, err := c.usecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peminjamans)
} 