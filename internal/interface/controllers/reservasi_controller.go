package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type ReservasiController struct {
	usecase *usecase.ReservasiUsecase
}

func NewReservasiController(usecase *usecase.ReservasiUsecase) *ReservasiController {
	return &ReservasiController{usecase: usecase}
}

func (c *ReservasiController) Create(w http.ResponseWriter, r *http.Request) {
	var reservasi entity.Reservasi
	if err := json.NewDecoder(r.Body).Decode(&reservasi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.usecase.Create(r.Context(), &reservasi); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservasi)
}

func (c *ReservasiController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var reservasi entity.Reservasi
	if err := json.NewDecoder(r.Body).Decode(&reservasi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reservasi.IDReservasi = id
	if err := c.usecase.Update(r.Context(), &reservasi); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservasi)
}

func (c *ReservasiController) Delete(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Reservasi berhasil dihapus"})
}

func (c *ReservasiController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	reservasi, err := c.usecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Reservasi tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservasi)
}

func (c *ReservasiController) GetAll(w http.ResponseWriter, r *http.Request) {
	reservasis, err := c.usecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservasis)
} 