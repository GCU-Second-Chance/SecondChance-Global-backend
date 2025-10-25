package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/config"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/model"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func GetDogsRandomFromGyeonggi(ctx context.Context) ([]*model.Dog, error) {
	apiKey := config.Cfg.Gyeonggi.GyeonggiApiKey
	randomPage := rand.Intn(GyeonggiPIndex) + 1
	baseURL := GyeonggiGetAnimalsURL

	params := url.Values{}
	params.Add("KEY", apiKey)
	params.Add("Type", "json")
	params.Add("pIndex", fmt.Sprintf("%d", randomPage))
	params.Add("pSize", fmt.Sprintf("%d", GyeonggiPSize))
	params.Add("STATE_NM", GyeonggiStatus)

	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Printf("📡 [요청] %s\n", reqURL)

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get gyeonggi animals: build request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get gyeonggi animals: send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("could not close response body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get gyeonggi animals: bad status code: %d", resp.StatusCode)
	}

	var payload model.GyeonggiRandomResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("failed to get petfinder random dogs: decode response: %w", err)
	}

	var dogs []*model.Dog
	for _, row := range payload.AbdmAnimalProtect[1].Row {
		dog := mapGyeonggiToDog(row)
		copiedDog := dog
		dogs = append(dogs, copiedDog)
	}
	return dogs, nil
}

func mapGyeonggiToDog(row model.AbdmRow) *model.Dog {
	id, _ := strconv.ParseInt(row.AbdmIDntfyNo, 10, 64)
	var images []string
	images = append(images, row.ThumbImageCours)

	return &model.Dog{
		ID:     id,
		Name:   GyeonggiEmpty,
		Age:    row.AgeInfo,
		Images: images,
		Gender: row.SexNM,
		Breed:  row.SpeciesNM,
		Location: model.Location{
			Country: GyeonggiCountry,
			City:    row.JurisdInstNM,
		},
		Shelter: model.Shelter{
			Name:    row.ShterNM,
			Contact: row.ShterTelno,
			Email:   GyeonggiEmpty,
		},
		CountryType: GyeonggiCountryType,
	}
}
