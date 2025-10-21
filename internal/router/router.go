package router

import (
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/handler"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	healthHandler *handler.HealthHandler
	dogHandler    *handler.DogHandler
}

func NewRouter() *Router {
	healthService := service.NewHealthService()
	dogService := service.NewDogService()

	healthHandler := handler.NewHealthHandler(healthService)
	dogHandler := handler.NewDogHandler(dogService)

	return &Router{
		healthHandler: healthHandler,
		dogHandler:    dogHandler,
	}
}

func (r *Router) SetupRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // TODO: 프론트엔드 주소로 변경
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/api/v1", func(router chi.Router) {
		// Health
		router.Get("/health", r.healthHandler.GetHealth)

		// Dog
		router.Route("/dogs", func(router chi.Router) {
			router.Get("/random", r.dogHandler.GetRandomDog)
			router.Get("/", r.dogHandler.GetDogByID)
		})
	})

	return router
}
