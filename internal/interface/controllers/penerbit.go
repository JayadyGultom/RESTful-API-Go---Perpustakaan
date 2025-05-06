package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type PenerbitController struct {
    usecase usecase.PenerbitUsecase
}

func NewPenerbitController(usecase usecase.PenerbitUsecase) *PenerbitController {
    return &PenerbitController{
        usecase: usecase,
    }
}

// HandlePenerbit: /penerbit endpoint (tanpa ID)
func (c *PenerbitController) HandlePenerbit(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        penerbit, err := c.usecase.GetAllPenerbit()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(penerbit)

    case http.MethodPost:
        var penerbit entity.Penerbit
        if err := json.NewDecoder(r.Body).Decode(&penerbit); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newPenerbit, err := c.usecase.CreatePenerbit(penerbit)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newPenerbit)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandlePenerbitByID: /penerbit/{id} endpoint
func (c *PenerbitController) HandlePenerbitByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/penerbit/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        penerbit, err := c.usecase.GetPenerbitByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(penerbit)

    case http.MethodPut:
        var penerbit entity.Penerbit
        if err := json.NewDecoder(r.Body).Decode(&penerbit); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdatePenerbit(id, penerbit)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeletePenerbit(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 