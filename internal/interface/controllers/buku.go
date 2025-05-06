package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type BukuController struct {
    usecase usecase.BukuUsecase
}

func NewBukuController(usecase usecase.BukuUsecase) *BukuController {
    return &BukuController{
        usecase: usecase,
    }
}

// HandleBuku: /buku endpoint (tanpa ID)
func (c *BukuController) HandleBuku(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        buku, err := c.usecase.GetAllBuku()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(buku)

    case http.MethodPost:
        var buku entity.Buku
        if err := json.NewDecoder(r.Body).Decode(&buku); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newBuku, err := c.usecase.CreateBuku(buku)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newBuku)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandleBukuByID: /buku/{id} endpoint
func (c *BukuController) HandleBukuByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/buku/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        buku, err := c.usecase.GetBukuByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(buku)

    case http.MethodPut:
        var buku entity.Buku
        if err := json.NewDecoder(r.Body).Decode(&buku); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdateBuku(id, buku)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeleteBuku(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 