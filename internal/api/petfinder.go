package api

import (
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
	TotalPages = 1715 // 최대 Page 범위
	Type       = "dog"
	Status     = "adoptable"
	Limit      = "10"
)

func GetDogByIDFromPetfinder(ctx context.Context, id int64) (*model.Dog, error) {
	url := util.PetfinderGetAnimalsURL + "/" + strconv.FormatInt(id, 10)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not request petfinder: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.Cfg.Petfinder.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not request petfinder: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request petfinder api: %s", resp.Status)
	}

	var payload model.PetfinderDogResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("could not decode petfinder random api response: %w", err)
	}

	dog := mapAnimalToDog(payload.Animal)
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
		return nil, fmt.Errorf("could not request petfinder: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.Cfg.Petfinder.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("could not close response body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request petfinder random api: %s", resp.Status)
	}

	var payload model.PetfinderDogsRandomResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("could not decode petfinder random api response: %w", err)
	}

	var dogs []*model.Dog
	for _, animal := range payload.Animals {
		dog := mapAnimalToDog(animal)
		copiedDog := dog
		dogs = append(dogs, copiedDog)
	}
	return dogs, nil
}

func mapAnimalToDog(animal model.Animal) *model.Dog {
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
			Name:    "Petfinder",
			Contact: animal.Contact.Phone,
			Email:   animal.Contact.Email,
		},
		CountryType: util.PetfinderCountryType,
	}
}
