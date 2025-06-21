package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type DendaController struct {
	usecase *usecase.DendaUsecase
}

func NewDendaController(usecase *usecase.DendaUsecase) *DendaController {
	return &DendaController{usecase: usecase}
}

func (c *DendaController) Create(w http.ResponseWriter, r *http.Request) {
	var denda entity.Denda
	if err := json.NewDecoder(r.Body).Decode(&denda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.usecase.Create(r.Context(), &denda); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(denda)
}

func (c *DendaController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var denda entity.Denda
	if err := json.NewDecoder(r.Body).Decode(&denda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	denda.IDDenda = id
	if err := c.usecase.Update(r.Context(), &denda); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(denda)
}

func (c *DendaController) Delete(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Denda berhasil dihapus"})
}

func (c *DendaController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	denda, err := c.usecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Denda tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(denda)
}

func (c *DendaController) GetAll(w http.ResponseWriter, r *http.Request) {
	denda, err := c.usecase.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(denda)
} 