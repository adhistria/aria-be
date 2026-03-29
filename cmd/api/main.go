package main

import (
	"log"
	"net/http"

	_ "github.com/adhistria/aria-be/docs" // swagger docs
	"github.com/adhistria/aria-be/internal/handler/rest"
	"github.com/adhistria/aria-be/internal/usecase/config"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Aria BE API
// @version         1.0
// @description     Aria Backend API
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	appConfigUsecase := config.NewAppConfigUsecase()
	appConfigHandler := rest.NewAppConfigHandler(appConfigUsecase)

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/app-config", appConfigHandler.GetAppConfig).Methods(http.MethodGet)

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
