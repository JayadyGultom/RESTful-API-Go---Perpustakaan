package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type PenulisController struct {
	penulisUsecase *usecase.PenulisUsecase
}

func NewPenulisController(penulisUsecase *usecase.PenulisUsecase) *PenulisController {
	return &PenulisController{penulisUsecase: penulisUsecase}
}

func (c *PenulisController) Create(w http.ResponseWriter, r *http.Request) {
	var penulis entity.Penulis
	if err := json.NewDecoder(r.Body).Decode(&penulis); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.penulisUsecase.Create(r.Context(), &penulis); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(penulis)
}

func (c *PenulisController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var penulis entity.Penulis
	if err := json.NewDecoder(r.Body).Decode(&penulis); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	penulis.IDPenulis = id
	if err := c.penulisUsecase.Update(r.Context(), &penulis); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penulis)
}

func (c *PenulisController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.penulisUsecase.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Penulis berhasil dihapus"})
}

func (c *PenulisController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	penulis, err := c.penulisUsecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Penulis tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penulis)
}

func (c *PenulisController) GetAll(w http.ResponseWriter, r *http.Request) {
	penulis, err := c.penulisUsecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penulis)
}

func (c *PenulisController) GetByNama(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("nama")
	if nama == "" {
		http.Error(w, "Nama penulis tidak boleh kosong", http.StatusBadRequest)
		return
	}

	penulis, err := c.penulisUsecase.GetByNama(r.Context(), nama)
	if err != nil {
		http.Error(w, "Penulis tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penulis)
} 