package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/config"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/model"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/util"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	TotalPages = 1715
	Type       = "dog"
	Status     = "adoptable"
	Limit      = "100"
)

func GetDogByID(ctx context.Context, id int64) (*model.Dog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.ContextTimeout)
	defer cancel()

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

	dog := mapToDog(payload.Animal)
	return dog, nil
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

func GetDogsRandom(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), util.ContextTimeout)
	defer cancel()

	randomPage := rand.Intn(TotalPages)

	url := util.PetfinderGetAnimalsURL +
		"?type=" + Type +
		"&status=" + Status +
		"&limit" + Limit +
		"&page=" + strconv.Itoa(randomPage)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not request petfinder: %w", err)
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
		return fmt.Errorf("failed to request petfinder random api: %s", resp.Status)
	}

	var payload model.PetfinderDogsRandomResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return fmt.Errorf("could not decode petfinder random api response: %w", err)
	}

	var dogs []*model.Dog
	for _, animal := range payload.Animals {
		dog := mapToDog(animal)
		copiedDog := dog
		dogs = append(dogs, copiedDog)
	}
	return nil
}
