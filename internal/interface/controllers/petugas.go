package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type PetugasController struct {
    usecase usecase.PetugasUsecase
}

func NewPetugasController(usecase usecase.PetugasUsecase) *PetugasController {
    return &PetugasController{
        usecase: usecase,
    }
}

// HandlePetugas: /petugas endpoint (tanpa ID)
func (c *PetugasController) HandlePetugas(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        petugas, err := c.usecase.GetAllPetugas()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(petugas)

    case http.MethodPost:
        var petugas entity.Petugas
        if err := json.NewDecoder(r.Body).Decode(&petugas); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newPetugas, err := c.usecase.CreatePetugas(petugas)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newPetugas)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandlePetugasByID: /petugas/{id} endpoint
func (c *PetugasController) HandlePetugasByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/petugas/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        petugas, err := c.usecase.GetPetugasByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(petugas)

    case http.MethodPut:
        var petugas entity.Petugas
        if err := json.NewDecoder(r.Body).Decode(&petugas); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdatePetugas(id, petugas)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeletePetugas(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 