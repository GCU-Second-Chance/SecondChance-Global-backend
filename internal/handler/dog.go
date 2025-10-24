package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/model"
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
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	countryStr := r.URL.Query().Get("country")
	if countryStr == "" {
		http.Error(w, "country query parameter is required", http.StatusBadRequest)
		return
	}

	country, err := model.StringToCountryType(countryStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := h.dogService.GetDogByIDJSON(ctx, country, idInt64)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get dog: %v", err), http.StatusInternalServerError)
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
		http.Error(w, fmt.Sprintf("failed to get random dog: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
