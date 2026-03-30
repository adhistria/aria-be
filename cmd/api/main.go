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

// corsMiddleware adds CORS headers to responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	appConfigUsecase := config.NewAppConfigUsecase()
	appConfigHandler := rest.NewAppConfigHandler(appConfigUsecase)

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/app-config", appConfigHandler.GetAppConfig).Methods(http.MethodGet)

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", corsMiddleware(r)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
