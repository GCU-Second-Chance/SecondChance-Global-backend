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
)

func GetTokenFromPetfinder(ctx context.Context) (string, error) {
	clientID := config.Cfg.Petfinder.ClientID
	clientSecret := config.Cfg.Petfinder.ClientSecret
	url := PetfinderGetAccessTokenURL
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

func GetDogByIDFromPetfinder(ctx context.Context, petfinderToken string, id int64) (*model.Dog, error) {
	url := PetfinderGetAnimalsURL + "/" + strconv.FormatInt(id, 10)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get petfinder dog %d: build request: %w", id, err)
	}

	req.Header.Set("Authorization", "Bearer "+petfinderToken)
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

	dog := mapPetfinerToDog(payload.Animal)
	return dog, nil
}

func GetDogsRandomFromPetfinder(ctx context.Context, petfinderToken string) ([]*model.Dog, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request was cancelled")
	default:
	}

	randomPage := rand.Intn(PetfinderTotalPages)

	url := PetfinderGetAnimalsURL +
		"?type=" + PetfinderType +
		"&status=" + PetfinderStatus +
		"&limit" + PetfinderLimit +
		"&page=" + strconv.Itoa(randomPage)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get petfinder random dogs: build request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+petfinderToken)

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
		dog := mapPetfinerToDog(animal)
		copiedDog := dog
		dogs = append(dogs, copiedDog)
	}
	return dogs, nil
}

func mapPetfinerToDog(animal model.Animal) *model.Dog {
	var images []string

	for _, photos := range animal.Photos {
		small := photos.Small
		medium := photos.Medium
		large := photos.Large
		full := photos.Full
		images = append(images, small, medium, large, full)
	}

	return &model.Dog{
		ID:     animal.ID,
		Name:   animal.Name,
		Age:    animal.Age,
		Images: images,
		Gender: animal.Gender,
		Breed:  animal.Breeds.Primary,
		Location: model.Location{
			Country: animal.Contact.Address.Country,
			City:    animal.Contact.Address.City,
		},
		Shelter: model.Shelter{
			Name:    PetfinderShelterName,
			Contact: animal.Contact.Phone,
			Email:   animal.Contact.Email,
		},
		CountryType: PetfinderCountryType,
	}
}
