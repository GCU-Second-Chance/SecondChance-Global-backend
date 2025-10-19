package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ChallengeService struct{}

type PreSignedUrlRequest struct {
	Prefix   string `json:"prefix"`
	FileName string `json:"filename"`
}

type PreSignedUrlResponse struct {
	Message       string    `json:"message"`
	PreSignedUrls []string  `json:"pre_signed_urls"`
	ExpiresAt     time.Time `json:"expires_at"`
	Status        string    `json:"status"`
}

func NewChallengeService() *ChallengeService {
	return &ChallengeService{}
}

func (s *ChallengeService) GetPreSignedUrl(ctx context.Context, request *PreSignedUrlRequest) *PreSignedUrlResponse {
	select {
	case <-ctx.Done():
		return &PreSignedUrlResponse{
			Message: "Request was cancelled",
			Status:  "cancelled",
		}
	default:
	}

	// TODO: S3에서 presigned URL 생성

	// 임시 데이터
	objectKey := fmt.Sprintf("%s/%s", request.Prefix, request.FileName)
	preSignedUrls := []string{
		fmt.Sprintf("https://s3.amazonaws.com/bucket/%s?presigned=...", objectKey),
	}

	expiresAt := time.Now().Add(1 * time.Hour)

	return &PreSignedUrlResponse{
		Message:       "Presigned URL generated successfully",
		PreSignedUrls: preSignedUrls,
		ExpiresAt:     expiresAt,
		Status:        "generated",
	}
}

func (s *ChallengeService) GetPreSignedUrlJSON(ctx context.Context, request *PreSignedUrlRequest) ([]byte, error) {
	response := s.GetPreSignedUrl(ctx, request)
	return json.Marshal(response)
}
