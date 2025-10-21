package handler

import (
	"net/http"

	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type DogHandler struct {
	dogService *service.DogService
}

func NewDogHandler(dogService *service.DogService) *DogHandler {
	return &DogHandler{
		dogService: dogService,
	}
}

func (h *DogHandler) GetDogByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	jsonData, err := h.dogService.GetDogByIDJSON(ctx, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *DogHandler) GetRandomDog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := h.dogService.GetRandomDogJSON(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
