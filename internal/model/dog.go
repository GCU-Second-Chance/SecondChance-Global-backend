package model

import "fmt"

type CountryType int

const (
	American CountryType = iota
	Korean
)

func StringToCountryType(countryStr string) (CountryType, error) {
	switch countryStr {
	case "American":
		return American, nil
	case "Korean":
		return Korean, nil
	default:
		return -1, fmt.Errorf("invalid country: %s", countryStr)
	}
}

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

type GetDogByIDRequest struct {
	Country string `json:"country"`
	ID      int64  `json:"id"`
}
