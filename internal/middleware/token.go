package middleware

import (
	"context"
	"net/http"

	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/api"
)

const PetfinderTokenKey = "petfinder_token"

func GetPetfinderToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := api.GetTokenFromPetfinder(r.Context())
		if err != nil {
			http.Error(w, "Failed to get Petfinder token", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), PetfinderTokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
