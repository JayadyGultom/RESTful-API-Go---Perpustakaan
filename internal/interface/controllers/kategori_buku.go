package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type KategoriBukuController struct {
    usecase usecase.KategoriBukuUsecase
}

func NewKategoriBukuController(usecase usecase.KategoriBukuUsecase) *KategoriBukuController {
    return &KategoriBukuController{
        usecase: usecase,
    }
}

// HandleKategoriBuku: /kategori-buku endpoint (tanpa ID)
func (c *KategoriBukuController) HandleKategoriBuku(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        kategori, err := c.usecase.GetAllKategoriBuku()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(kategori)

    case http.MethodPost:
        var kategori entity.KategoriBuku
        if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newKategori, err := c.usecase.CreateKategoriBuku(kategori)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newKategori)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandleKategoriBukuByID: /kategori-buku/{id} endpoint
func (c *KategoriBukuController) HandleKategoriBukuByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/kategori-buku/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        kategori, err := c.usecase.GetKategoriBukuByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(kategori)

    case http.MethodPut:
        var kategori entity.KategoriBuku
        if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdateKategoriBuku(id, kategori)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeleteKategoriBuku(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 