package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/api"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/middleware"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/model"
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

	token, ok := ctx.Value(middleware.PetfinderTokenKey).(string)
	if !ok {
		return nil, fmt.Errorf("petfinder token not found in context")
	}

	switch country {
	case model.American:
		americanDog, err := api.GetDogByIDFromPetfinder(ctx, token, id)
		if err != nil {
			return nil, err
		}
		return &model.DogResponse{
			Message: fmt.Sprintf("Get American dog with id: %d", id),
			Data:    americanDog,
		}, nil
	case model.Korean:
		koreanDog, err := api.GetDogByIDFromPetfinder(ctx, token, id)
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
	token, ok := ctx.Value(middleware.PetfinderTokenKey).(string)
	if !ok {
		return nil, fmt.Errorf("petfinder token not found in context")
	}

	americanDogsData, err := api.GetDogsRandomFromPetfinder(ctx, token)
	if err != nil {
		return nil, err
	}
	koreanDogsData, err := api.GetDogsRandomFromGyeonggi(ctx)
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
