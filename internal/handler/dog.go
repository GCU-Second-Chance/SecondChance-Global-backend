package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/model"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/service"
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

	var request model.GetDogByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	country, err := model.StringToCountryType(request.Country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := h.dogService.GetDogByIDJSON(ctx, country, request.ID)
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
