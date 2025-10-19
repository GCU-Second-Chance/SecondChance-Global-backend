package service

import (
	"context"
	"encoding/json"
	"time"
)

type HealthService struct{}

type HealthResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) GetHealth(ctx context.Context) *HealthResponse {
	// Context 취소 확인
	select {
	case <-ctx.Done():
		return &HealthResponse{
			Status:    "cancelled",
			Message:   "Request was cancelled",
			Timestamp: time.Now(),
		}
	default:
		// 정상 처리
	}

	return &HealthResponse{
		Status:    "ok",
		Message:   "Server is running",
		Timestamp: time.Now(),
	}
}

func (s *HealthService) GetHealthJSON(ctx context.Context) ([]byte, error) {
	health := s.GetHealth(ctx)
	return json.Marshal(health)
}
