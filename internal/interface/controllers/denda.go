package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "perpustakaan/internal/entity"
    "perpustakaan/internal/usecase"
)

type DendaController struct {
    usecase usecase.DendaUsecase
}

func NewDendaController(usecase usecase.DendaUsecase) *DendaController {
    return &DendaController{
        usecase: usecase,
    }
}

// HandleDenda: /denda endpoint (tanpa ID)
func (c *DendaController) HandleDenda(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        denda, err := c.usecase.GetAllDenda()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(denda)

    case http.MethodPost:
        var denda entity.Denda
        if err := json.NewDecoder(r.Body).Decode(&denda); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        newDenda, err := c.usecase.CreateDenda(denda)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newDenda)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// HandleDendaByID: /denda/{id} endpoint
func (c *DendaController) HandleDendaByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/denda/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        denda, err := c.usecase.GetDendaByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(denda)

    case http.MethodPut:
        var denda entity.Denda
        if err := json.NewDecoder(r.Body).Decode(&denda); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        updated, err := c.usecase.UpdateDenda(id, denda)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(updated)

    case http.MethodDelete:
        err := c.usecase.DeleteDenda(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
} 