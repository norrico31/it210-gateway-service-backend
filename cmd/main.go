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
	authHandler := proxy.ReverseProxy(config.Envs.AuthURL)

	coreHandler := proxy.ReverseProxy(config.Envs.CoreURL)

	// PUBLIC ROUTES
	router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath+"/login"), authHandler)

	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/users"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/users/deleted"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/users/{userId}"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/users/{userId}/restore"), coreHandler)

	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/roles"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/roles/{roleId}"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/roles/{roleId}/restore"), coreHandler)

	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/projects"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/projects/deleted"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/projects/{projectId}"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/projects/{projectId}/restore"), coreHandler)

	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/tasks"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/tasks/deleted"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/tasks/{taskId}"), coreHandler)
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/tasks/{taskId}/restore"), coreHandler)

	// AUTH ROUTES
	// router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath+"/login"), middleware.ValidateJWT(authHandler))

	// CORE ROUTES
	// router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath+"/login"), authHandler)
	// Add more routes as needed

	log.Println("Gateway is running on port 8083")
	if err := http.ListenAndServe(":8083", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
