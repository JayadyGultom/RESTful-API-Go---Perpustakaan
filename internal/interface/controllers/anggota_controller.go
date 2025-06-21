package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
)

type AnggotaController struct {
	Usecase usecase.AnggotaUsecase
}

func (c *AnggotaController) GetAll(w http.ResponseWriter, r *http.Request) {
	anggotas, err := c.Usecase.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusBadRequest)
json.NewEncoder(w).Encode(map[string]string{
    "error": err.Error(),
})
		return
	}
	json.NewEncoder(w).Encode(anggotas)
}

func (c *AnggotaController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	anggota, err := c.Usecase.GetByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusBadRequest)
json.NewEncoder(w).Encode(map[string]string{
    "error": err.Error(),
})
		return
	}
	json.NewEncoder(w).Encode(anggota)
}

func (c *AnggotaController) Create(w http.ResponseWriter, r *http.Request) {
	var a entity.Anggota
	json.NewDecoder(r.Body).Decode(&a)
	err := c.Usecase.Create(a)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *AnggotaController) Update(w http.ResponseWriter, r *http.Request) {
	var a entity.Anggota
	json.NewDecoder(r.Body).Decode(&a)
	err := c.Usecase.Update(a)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusBadRequest)
json.NewEncoder(w).Encode(map[string]string{
    "error": err.Error(),
})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *AnggotaController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	err := c.Usecase.Delete(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusBadRequest)
json.NewEncoder(w).Encode(map[string]string{
    "error": err.Error(),
})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewAnggotaController(usecase usecase.AnggotaUsecase) *AnggotaController {
	return &AnggotaController{Usecase: usecase}
}
