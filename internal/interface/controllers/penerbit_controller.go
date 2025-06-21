package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type PenerbitController struct {
	penerbitUsecase *usecase.PenerbitUsecase
}

func NewPenerbitController(penerbitUsecase *usecase.PenerbitUsecase) *PenerbitController {
	return &PenerbitController{penerbitUsecase: penerbitUsecase}
}

func (c *PenerbitController) Create(w http.ResponseWriter, r *http.Request) {
	var penerbit entity.Penerbit
	if err := json.NewDecoder(r.Body).Decode(&penerbit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.penerbitUsecase.Create(r.Context(), &penerbit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(penerbit)
}

func (c *PenerbitController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var penerbit entity.Penerbit
	if err := json.NewDecoder(r.Body).Decode(&penerbit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	penerbit.IDPenerbit = id
	if err := c.penerbitUsecase.Update(r.Context(), &penerbit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penerbit)
}

func (c *PenerbitController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.penerbitUsecase.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Penerbit berhasil dihapus"})
}

func (c *PenerbitController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	penerbit, err := c.penerbitUsecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Penerbit tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penerbit)
}

func (c *PenerbitController) GetAll(w http.ResponseWriter, r *http.Request) {
	penerbit, err := c.penerbitUsecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penerbit)
}

func (c *PenerbitController) GetByNama(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("nama")
	if nama == "" {
		http.Error(w, "Nama penerbit tidak boleh kosong", http.StatusBadRequest)
		return
	}

	penerbit, err := c.penerbitUsecase.GetByNama(r.Context(), nama)
	if err != nil {
		http.Error(w, "Penerbit tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penerbit)
} 