package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type PeminjamanController struct {
    usecase usecase.PeminjamanUsecase
}

func NewPeminjamanController(usecase usecase.PeminjamanUsecase) *PeminjamanController {
    return &PeminjamanController{
        usecase: usecase,
    }
}

// HandlePeminjaman: /peminjaman endpoint (tanpa ID)
func (c *PeminjamanController) HandlePeminjaman(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        peminjaman, err := c.usecase.GetAllPeminjaman()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(peminjaman)

    case http.MethodPost:
        var peminjaman entity.Peminjaman
        if err := json.NewDecoder(r.Body).Decode(&peminjaman); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newPeminjaman, err := c.usecase.CreatePeminjaman(peminjaman)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newPeminjaman)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandlePeminjamanByID: /peminjaman/{id} endpoint
func (c *PeminjamanController) HandlePeminjamanByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/peminjaman/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        peminjaman, err := c.usecase.GetPeminjamanByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(peminjaman)

    case http.MethodPut:
        var peminjaman entity.Peminjaman
        if err := json.NewDecoder(r.Body).Decode(&peminjaman); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdatePeminjaman(id, peminjaman)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeletePeminjaman(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 