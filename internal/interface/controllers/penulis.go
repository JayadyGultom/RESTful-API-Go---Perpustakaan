package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type PenulisController struct {
    usecase usecase.PenulisUsecase
}

func NewPenulisController(usecase usecase.PenulisUsecase) *PenulisController {
    return &PenulisController{
        usecase: usecase,
    }
}

// HandlePenulis: /penulis endpoint (tanpa ID)
func (c *PenulisController) HandlePenulis(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        penulis, err := c.usecase.GetAllPenulis()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(penulis)

    case http.MethodPost:
        var penulis entity.Penulis
        if err := json.NewDecoder(r.Body).Decode(&penulis); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newPenulis, err := c.usecase.CreatePenulis(penulis)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newPenulis)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandlePenulisByID: /penulis/{id} endpoint
func (c *PenulisController) HandlePenulisByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/penulis/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        penulis, err := c.usecase.GetPenulisByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(penulis)

    case http.MethodPut:
        var penulis entity.Penulis
        if err := json.NewDecoder(r.Body).Decode(&penulis); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdatePenulis(id, penulis)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeletePenulis(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 