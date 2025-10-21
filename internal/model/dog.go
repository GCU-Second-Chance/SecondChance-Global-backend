package model

import "fmt"

type CountryType int

const (
	American CountryType = iota
	Korean
)

type Dog struct {
	ID          int64
	Name        string
	Age         string
	Images      []string
	Gender      string
	Breed       string
	Location    Location
	Shelter     Shelter
	CountryType string
}

type Location struct {
	Country string
	City    string
}

type Shelter struct {
	Name    string
	Contact string
	Email   string
}

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

type DogsResponse struct {
	Message string `json:"message"`
	Data    []*Dog `json:"data,omitempty"`
}

type DogResponse struct {
	Message string `json:"message"`
	Data    *Dog   `json:"data,omitempty"`
}

type GetDogByIDRequest struct {
	Country string `json:"country"`
	ID      int64  `json:"id"`
}
