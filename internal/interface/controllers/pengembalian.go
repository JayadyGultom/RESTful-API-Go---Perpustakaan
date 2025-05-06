package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type PengembalianController struct {
    usecase usecase.PengembalianUsecase
}

func NewPengembalianController(usecase usecase.PengembalianUsecase) *PengembalianController {
    return &PengembalianController{
        usecase: usecase,
    }
}

// HandlePengembalian: /pengembalian endpoint (tanpa ID)
func (c *PengembalianController) HandlePengembalian(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        pengembalian, err := c.usecase.GetAllPengembalian()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(pengembalian)

    case http.MethodPost:
        var pengembalian entity.Pengembalian
        if err := json.NewDecoder(r.Body).Decode(&pengembalian); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newPengembalian, err := c.usecase.CreatePengembalian(pengembalian)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newPengembalian)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandlePengembalianByID: /pengembalian/{id} endpoint
func (c *PengembalianController) HandlePengembalianByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/pengembalian/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        pengembalian, err := c.usecase.GetPengembalianByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(pengembalian)

    case http.MethodPut:
        var pengembalian entity.Pengembalian
        if err := json.NewDecoder(r.Body).Decode(&pengembalian); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdatePengembalian(id, pengembalian)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeletePengembalian(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 