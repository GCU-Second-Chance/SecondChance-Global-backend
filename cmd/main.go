package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/config"
	"github.com/GCU-Second-Chance/SecondChance-Global-backend/internal/router"
)

func main() {
	config.Load()

	routerInstance := router.NewRouter()
	r := routerInstance.SetupRoutes()

	serverAddr := fmt.Sprintf("%s:%s", config.Cfg.Server.Host, config.Cfg.Server.Port)

	log.Printf("Server starting on %s", serverAddr)
	log.Printf("Health check: http://%s/api/health", serverAddr)
	log.Printf("Available endpoints:")
	log.Printf("  GET  /api/health")
	log.Printf("  GET  /api/dogs/random")
	log.Printf("  GET  /api/dogs/{id}")
	log.Printf("  POST /api/challenge/upload")

	if err := http.ListenAndServe(serverAddr, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
