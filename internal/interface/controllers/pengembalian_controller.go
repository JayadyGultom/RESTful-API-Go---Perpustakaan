package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type PengembalianController struct {
	usecase *usecase.PengembalianUsecase
}

func NewPengembalianController(usecase *usecase.PengembalianUsecase) *PengembalianController {
	return &PengembalianController{usecase: usecase}
}

func (c *PengembalianController) Create(w http.ResponseWriter, r *http.Request) {
	var pengembalian entity.Pengembalian
	if err := json.NewDecoder(r.Body).Decode(&pengembalian); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.usecase.Create(r.Context(), &pengembalian); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pengembalian)
}

func (c *PengembalianController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var pengembalian entity.Pengembalian
	if err := json.NewDecoder(r.Body).Decode(&pengembalian); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pengembalian.IDPengembalian = id
	if err := c.usecase.Update(r.Context(), &pengembalian); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pengembalian)
}

func (c *PengembalianController) Delete(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Pengembalian berhasil dihapus"})
}

func (c *PengembalianController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	pengembalian, err := c.usecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Pengembalian tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pengembalian)
}

func (c *PengembalianController) GetAll(w http.ResponseWriter, r *http.Request) {
	pengembalians, err := c.usecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pengembalians)
} 