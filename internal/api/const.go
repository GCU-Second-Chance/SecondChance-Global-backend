package api

const (
	GyeonggiGetAnimalsURL = "https://openapi.gg.go.kr/AbdmAnimalProtect"
	GyeonggiStatus        = "보호중"
	GyeonggiTotalCount    = 2129
	GyeonggiPSize         = 10
	GyeonggiPIndex        = int(GyeonggiTotalCount / GyeonggiPSize)
	GyeonggiEmpty         = "빈값"
	GyeonggiCountry       = "대한민국"
	GyeonggiCountryType   = "Korean"

	PetfinderGetAccessTokenURL = "https://api.petfinder.com/v2/oauth2/token"
	PetfinderGetAnimalsURL     = "https://api.petfinder.com/v2/animals"
	PetfinderTotalPages        = 1715
	PetfinderType              = "dog"
	PetfinderStatus            = "adoptable"
	PetfinderLimit             = "10"
	PetfinderShelterName       = "Petfinder"
	PetfinderCountryType       = "American"
)
