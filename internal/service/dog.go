package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"SecondChance-Global-backend/internal/model"
)

type DogService struct{}

func NewDogService() *DogService {
	return &DogService{}
}

func (s *DogService) GetDogByID(ctx context.Context, country model.CountryType, id int64) (*model.DogResponse, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request was cancelled")
	default:
	}

	switch country {
	case model.American:
		americanDog, err := s.getAmericanDogDataById(id)
		if err != nil {
			return nil, err
		}
		return &model.DogResponse{
			Message: fmt.Sprintf("Get American dog with id: %d", id),
			Data:    americanDog,
		}, nil
	case model.Korean:
		koreanDog, err := s.getKoreanDogDataById(id)
		if err != nil {
			return nil, err
		}
		return &model.DogResponse{
			Message: fmt.Sprintf("Get Korean dog with id: %d", id),
			Data:    koreanDog,
		}, nil
	default:
		return nil, fmt.Errorf("invalid country: %v", country)
	}
}

func (s *DogService) GetRandomDog(ctx context.Context) (*model.DogsResponse, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request was cancelled")
	default:
	}

	americanDogsData, err := s.getAmericanDogsData()
	if err != nil {
		return nil, err
	}
	koreanDogsData, err := s.getKoreanDogsData()
	if err != nil {
		return nil, err
	}

	allDogsData := append(americanDogsData, koreanDogsData...)

	// 랜덤으로 섞기
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allDogsData), func(i, j int) {
		allDogsData[i], allDogsData[j] = allDogsData[j], allDogsData[i]
	})

	return &model.DogsResponse{
		Message: "Random dogs selected from American and Korean data",
		Data:    allDogsData,
	}, nil
}

func (s *DogService) getAmericanDogsData() ([]model.Dog, error) {
	return []model.Dog{
		{
			ID:          1,
			Name:        "mungmung",
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
			Name:        "bbobbi",
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
			Name:        "choco",
			Breed:       "시바견",
			Age:         4,
			Gender:      "수컷",
			Description: "독립적이지만 충성심이 강한 강아지입니다.",
			ImageURL:    "https://example.com/dog3.jpg",
			Location:    "서울시 마포구",
			Address:     "서울시 마포구 마포대로 789",
			Status:      "reserved",
		},
	}, nil
}

func (s *DogService) getAmericanDogDataById(id int64) (model.Dog, error) {
	switch id {
	case 1:
		return model.Dog{
			ID:          1,
			Name:        "mungmung",
			Breed:       "믹스",
			Age:         3,
			Gender:      "수컷",
			Description: "활발하고 친근한 성격의 강아지입니다.",
			ImageURL:    "https://example.com/dog1.jpg",
			Location:    "서울시 강남구",
			Address:     "서울시 강남구 강남대로 123",
			Status:      "available",
		}, nil
	case 2:
		return model.Dog{
			ID:          2,
			Name:        "bbobbi",
			Breed:       "골든리트리버",
			Age:         2,
			Gender:      "암컷",
			Description: "온순하고 똑똑한 강아지입니다.",
			ImageURL:    "https://example.com/dog2.jpg",
			Location:    "서울시 서초구",
			Address:     "서울시 서초구 서초대로 456",
			Status:      "available",
		}, nil
	case 3:
		return model.Dog{
			ID:          3,
			Name:        "choco",
			Breed:       "시바견",
			Age:         4,
			Gender:      "수컷",
			Description: "독립적이지만 충성심이 강한 강아지입니다.",
			ImageURL:    "https://example.com/dog3.jpg",
			Location:    "서울시 마포구",
			Address:     "서울시 마포구 마포대로 789",
			Status:      "reserved",
		}, nil
	default:
		return model.Dog{}, fmt.Errorf("invalid id: %d", id)
	}
}

func (s *DogService) getKoreanDogsData() ([]model.Dog, error) {
	return []model.Dog{
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
	}, nil
}

func (s *DogService) getKoreanDogDataById(id int64) (model.Dog, error) {
	switch id {
	case 1:
		return model.Dog{
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
		}, nil
	case 2:
		return model.Dog{
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
		}, nil
	case 3:
		return model.Dog{
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
		}, nil
	default:
		return model.Dog{}, fmt.Errorf("invalid id: %d", id)
	}
}

func (s *DogService) GetDogByIDJSON(ctx context.Context, country model.CountryType, id int64) ([]byte, error) {
	response, err := s.GetDogByID(ctx, country, id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(response)
}

func (s *DogService) GetRandomDogJSON(ctx context.Context) ([]byte, error) {
	response, err := s.GetRandomDog(ctx)
	if err != nil {
		return nil, err
	}
	return json.Marshal(response)
}
