package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type DetailPeminjamanController struct {
    usecase usecase.DetailPeminjamanUsecase
}

func NewDetailPeminjamanController(usecase usecase.DetailPeminjamanUsecase) *DetailPeminjamanController {
    return &DetailPeminjamanController{
        usecase: usecase,
    }
}

// HandleDetailPeminjaman: /detail-peminjaman endpoint (tanpa ID)
func (c *DetailPeminjamanController) HandleDetailPeminjaman(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        detail, err := c.usecase.GetAllDetailPeminjaman()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(detail)

    case http.MethodPost:
        var detail entity.DetailPeminjaman
        if err := json.NewDecoder(r.Body).Decode(&detail); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newDetail, err := c.usecase.CreateDetailPeminjaman(detail)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newDetail)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandleDetailPeminjamanByID: /detail-peminjaman/{id} endpoint
func (c *DetailPeminjamanController) HandleDetailPeminjamanByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/detail-peminjaman/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        detail, err := c.usecase.GetDetailPeminjamanByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(detail)

    case http.MethodPut:
        var detail entity.DetailPeminjaman
        if err := json.NewDecoder(r.Body).Decode(&detail); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdateDetailPeminjaman(id, detail)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeleteDetailPeminjaman(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 