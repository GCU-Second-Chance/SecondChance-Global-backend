package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

type DogService struct{}

type Dog struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	Status      string `json:"status"`
}

type DogsResponse struct {
	Message string `json:"message"`
	Data    []Dog  `json:"data,omitempty"`
}

type DogResponse struct {
	Message string `json:"message"`
	Data    Dog    `json:"data,omitempty"`
}

func NewDogService() *DogService {
	return &DogService{}
}

func (s *DogService) GetDogByID(ctx context.Context, id string) *DogResponse {
	select {
	case <-ctx.Done():
		return &DogResponse{
			Message: "Request was cancelled",
		}
	default:
	}

	dogID, err := strconv.Atoi(id)
	if err != nil {
		return &DogResponse{
			Message: "Invalid dog ID",
		}
	}

	// TODO: ID로 해당 유기견 조회 (외부 API 호출)

	// 임시 데이터에서 찾기
	if dogID < 1 || dogID > len(s.getDogsData()) {
		return &DogResponse{
			Message: "Dog not found",
		}
	}
	dogs := s.getDogsData()
	return &DogResponse{
		Message: fmt.Sprintf("Get dog with id: %d", dogID),
		Data:    dogs[dogID-1],
	}
}

func (s *DogService) GetRandomDog(ctx context.Context) *DogsResponse {
	select {
	case <-ctx.Done():
		return &DogsResponse{
			Message: "Request was cancelled",
		}
	default:
	}

	// TODO: 유기견 리스트 랜덤 조회 (외부 API 호출)

	dogs := s.getDogsData()

	return &DogsResponse{
		Message: "Random dogs selected",
		Data:    dogs,
	}
}

func (s *DogService) getDogsData() []Dog {
	return []Dog{
		{
			ID:          1,
			Name:        "멍멍이",
			Breed:       "믹스",
			Age:         3,
			Gender:      "수컷",
			Description: "활발하고 친근한 성격의 강아지입니다.",
			ImageURL:    "https://example.com/dog1.jpg",
			Location:    "서울시 강남구",
			Address:     "서울시 강남구 강남대로 123",
			Status:      "available",
		},
		{
			ID:          2,
			Name:        "뽀삐",
			Breed:       "골든리트리버",
			Age:         2,
			Gender:      "암컷",
			Description: "온순하고 똑똑한 강아지입니다.",
			ImageURL:    "https://example.com/dog2.jpg",
			Location:    "서울시 서초구",
			Address:     "서울시 서초구 서초대로 456",
			Status:      "available",
		},
		{
			ID:          3,
			Name:        "초코",
			Breed:       "시바견",
			Age:         4,
			Gender:      "수컷",
			Description: "독립적이지만 충성심이 강한 강아지입니다.",
			ImageURL:    "https://example.com/dog3.jpg",
			Location:    "서울시 마포구",
			Address:     "서울시 마포구 마포대로 789",
			Status:      "reserved",
		},
	}
}

func (s *DogService) GetDogByIDJSON(ctx context.Context, id string) ([]byte, error) {
	response := s.GetDogByID(ctx, id)
	return json.Marshal(response)
}

func (s *DogService) GetRandomDogJSON(ctx context.Context) ([]byte, error) {
	response := s.GetRandomDog(ctx)
	return json.Marshal(response)
}
