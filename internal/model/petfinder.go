package model

type PetfinderDogsRandomResponse struct {
	Animals []Animal
}

type PetfinderDogResponse struct {
	Animal Animal
}

type Animal struct {
	ID     int64
	Name   string
	Age    string
	Photos []struct {
		Small  string
		Medium string
		Large  string
		Full   string
	}

	Gender string

	Breeds struct {
		Primary string
	}

	PrimaryPhotoCropped struct {
		Small  string
		Medium string
		Large  string
		Full   string
	}
	Status          string
	StatusChangedAt string
	PublishedAt     string
	Contact         struct {
		Email   string
		Phone   string
		Address struct {
			Address1 string
			Address2 string
			City     string
			State    string
			Postcode string
			Country  string
		}
	}

	URL    string
	Colors struct {
		Primary string
	}
}
