package model

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
