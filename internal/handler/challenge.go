package handler

import (
	"encoding/json"
	"net/http"

	"SecondChance-Global-backend/internal/service"
)

type ChallengeHandler struct {
	challengeService *service.ChallengeService
}

func NewChallengeHandler(challengeService *service.ChallengeService) *ChallengeHandler {
	return &ChallengeHandler{
		challengeService: challengeService,
	}
}

func (h *ChallengeHandler) GetPreSignedUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json")

	// 요청 바디에서 prefix와 filename 파싱
	var request service.PreSignedUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	jsonData, err := h.challengeService.GetPreSignedUrlJSON(ctx, &request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
