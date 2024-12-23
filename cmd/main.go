package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

	router.HandleFunc("/api/v1/helloworld", http.HandlerFunc(helloWorldHandler), "GET")

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

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                                       // Allow all origins, or specify specific domains here
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allow specific methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Allow specific headers
	)(router)

	server := &http.Server{
		Addr:           ":8083",
		Handler:        corsHandler,
		MaxHeaderBytes: 1 << 20, // 1 MB for header size, adjust as needed
	}

	log.Println("Gateway is running on port 8083")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World! Gateway is working.")
}
