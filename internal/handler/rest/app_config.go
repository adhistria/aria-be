package rest

import (
	"encoding/json"
	"net/http"

	"github.com/adhistria/aria-be/internal/entity"
)

type AppConfigService interface {
	GetAppConfig() (*entity.AppConfig, error)
}

type AppConfigHandler struct {
	appConfigService AppConfigService
}

func NewAppConfigHandler(appConfigService AppConfigService) *AppConfigHandler {
	return &AppConfigHandler{
		appConfigService: appConfigService,
	}
}

// GetAppConfig godoc
// @Summary      Get application configuration
// @Description  Returns the current app configuration including version info and feature flags
// @Tags         app-config
// @Produce      json
// @Success      200  {object}  entity.AppConfig
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /app-config [get]
func (h *AppConfigHandler) GetAppConfig(w http.ResponseWriter, r *http.Request) {
	appConfig, err := h.appConfigService.GetAppConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appConfig)
}
