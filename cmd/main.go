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

	// PUBLIC ROUTES
	router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath), proxy.ReverseProxy("8081"))
	router.HandleFunc(utils.HandlePathV1(config.Envs.AuthPath+"/helloworld"), proxy.ReverseProxy("8081"))
	router.HandleFunc(utils.HandlePathV1(config.Envs.CorePath+"/helloworld"), proxy.ReverseProxy("8082"))

	// AUTH ROUTES
	router.HandleFunc("/api/v1/", middleware.ValidateJWT(proxy.ReverseProxy("8081")))

	log.Println("Gateway is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
