package main

import (
	"log"
	"net/http"

	"github.com/norrico31/it210-gateway-service-backend/config"
	"github.com/norrico31/it210-gateway-service-backend/internal/middleware"
	"github.com/norrico31/it210-gateway-service-backend/internal/proxy"
	"github.com/norrico31/it210-gateway-service-backend/internal/utils"
)

func main() {
	router := middleware.NewServeMux()

	// Create a reverse proxy handler for the auth service
	authHandler := proxy.ReverseProxy("8081")

	coreHandler := proxy.ReverseProxy("8082")

	// PUBLIC ROUTES
	router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath+"/login"), authHandler)

	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/users"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/roles"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/roles/{roleId}"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/roles/{roleId}/restore"), coreHandler)

	// AUTH ROUTES
	// router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath+"/login"), middleware.ValidateJWT(authHandler))

	// CORE ROUTES
	// router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath+"/login"), authHandler)
	// Add more routes as needed
	log.Println("Gateway is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
