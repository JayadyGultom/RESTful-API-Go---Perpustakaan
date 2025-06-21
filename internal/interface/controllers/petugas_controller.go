package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type PetugasController struct {
	usecase *usecase.PetugasUsecase
}

func NewPetugasController(usecase *usecase.PetugasUsecase) *PetugasController {
	return &PetugasController{usecase: usecase}
}

func (c *PetugasController) Create(w http.ResponseWriter, r *http.Request) {
	var petugas entity.Petugas
	if err := json.NewDecoder(r.Body).Decode(&petugas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.usecase.Create(r.Context(), &petugas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(petugas)
}

func (c *PetugasController) GetAll(w http.ResponseWriter, r *http.Request) {
	petugas, err := c.usecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petugas)
}

func (c *PetugasController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	petugas, err := c.usecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Petugas tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petugas)
}

func (c *PetugasController) Update(w http.ResponseWriter, r *http.Request) {
	var petugas entity.Petugas
	if err := json.NewDecoder(r.Body).Decode(&petugas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.usecase.Update(r.Context(), &petugas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petugas)
}

func (c *PetugasController) Delete(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Petugas berhasil dihapus"})
}

func (c *PetugasController) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username tidak boleh kosong", http.StatusBadRequest)
		return
	}
	petugas, err := c.usecase.GetByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, "Petugas tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petugas)
} 