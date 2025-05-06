package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
)

type AnggotaController struct {
	usecase usecase.AnggotaUsecase
}

func NewAnggotaController(usecase usecase.AnggotaUsecase) *AnggotaController {
	return &AnggotaController{
		usecase: usecase,
	}
}

// HandleAnggota: /anggota endpoint (tanpa ID)
func (c *AnggotaController) HandleAnggota(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		anggota, err := c.usecase.GetAllAnggota()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(anggota)

	case http.MethodPost:
		var anggota entity.Anggota
		if err := json.NewDecoder(r.Body).Decode(&anggota); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		newAnggota, err := c.usecase.CreateAnggota(anggota)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newAnggota)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// HandleAnggotaByID: /anggota/{id} endpoint
func (c *AnggotaController) HandleAnggotaByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/anggota/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		anggota, err := c.usecase.GetAnggotaByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(anggota)

	case http.MethodPut:
		var anggota entity.Anggota
		if err := json.NewDecoder(r.Body).Decode(&anggota); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		updated, err := c.usecase.UpdateAnggota(id, anggota)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		err := c.usecase.DeleteAnggota(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
