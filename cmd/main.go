package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/norrico31/it210-gateway-service-backend/internal/middleware"
	"github.com/norrico31/it210-gateway-service-backend/internal/proxy"
)

func main() {
	router := http.NewServeMux()
	godotenv.Load()

	// PUBLIC ROUTES
	router.Handle("/api/v1/users/helloworld", proxy.ReverseProxy("8081"))
	router.Handle("/api/v1/users", proxy.ReverseProxy("8081"))

	// AUTH ROUTES
	router.Handle("/api/v1/", middleware.ValidateJWT(proxy.ReverseProxy("8081")))

	log.Println("Gateway is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
