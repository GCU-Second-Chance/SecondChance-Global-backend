package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/config"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/model"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/util"
)

const (
	TotalPages = 1715
	Type       = "dog"
	Status     = "adoptable"
	Limit      = "10"
)

func GetTokenFromPetfinder(ctx context.Context) (string, error) {
	clientID := config.Cfg.Petfinder.ClientID
	clientSecret := config.Cfg.Petfinder.ClientSecret
	url := util.PetfinderGetAccessTokenURL
	data := []byte(fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret))

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("failed to get petfinder token: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get petfinder token: send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("could not close response body")
		}
	}(resp.Body)

	var tokenResp model.PetfinderTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to get petfinder token: decode response: %w", err)
	}
	return tokenResp.AccessToken, nil
}

func GetDogByIDFromPetfinder(ctx context.Context, id int64) (*model.Dog, error) {
	url := util.PetfinderGetAnimalsURL + "/" + strconv.FormatInt(id, 10)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get petfinder dog %d: build request: %w", id, err)
	}

	req.Header.Set("Authorization", "Bearer "+config.Cfg.Petfinder.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get petfinder dog %d: send request: %w", id, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("could not close response body")
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get petfinder dog %d: unexpected status %s", id, resp.Status)
	}

	var payload model.PetfinderDogResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("failed to get petfinder dog %d: decode response: %w", id, err)
	}

	dog := mapToDog(payload.Animal)
	return dog, nil
}

func GetDogsRandomFromPetfinder(ctx context.Context) ([]*model.Dog, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request was cancelled")
	default:
	}

	randomPage := rand.Intn(TotalPages)

	url := util.PetfinderGetAnimalsURL +
		"?type=" + Type +
		"&status=" + Status +
		"&limit" + Limit +
		"&page=" + strconv.Itoa(randomPage)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get petfinder random dogs: build request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.Cfg.Petfinder.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get petfinder random dogs: send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("could not close response body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get petfinder random dogs: unexpected status %s", resp.Status)
	}

	var payload model.PetfinderDogsRandomResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("failed to get petfinder random dogs: decode response: %w", err)
	}

	var dogs []*model.Dog
	for _, animal := range payload.Animals {
		dog := mapToDog(animal)
		copiedDog := dog
		dogs = append(dogs, copiedDog)
	}
	return dogs, nil
}

func mapToDog(Animal model.Animal) *model.Dog {
	var images []string

	for _, photos := range Animal.Photos {
		small := photos.Small
		medium := photos.Medium
		large := photos.Large
		full := photos.Full
		images = append(images, small, medium, large, full)
	}

	return &model.Dog{
		ID:     Animal.ID,
		Name:   Animal.Name,
		Age:    Animal.Age,
		Images: images,
		Gender: Animal.Gender,
		Breed:  Animal.Breeds.Primary,
		Location: model.Location{
			Country: Animal.Contact.Address.Country,
			City:    Animal.Contact.Address.City,
		},
		Shelter: model.Shelter{
			Name:    "Petfinder",
			Contact: Animal.Contact.Phone,
			Email:   Animal.Contact.Email,
		},
		CountryType: util.PetfinderCountryType,
	}
}
