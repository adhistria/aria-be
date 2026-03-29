package config

import (
	"github.com/adhistria/aria-be/internal/entity"
)

type appConfigUsecase struct {
}

func (a *appConfigUsecase) GetAppConfig() (*entity.AppConfig, error) {
	// In a real implementation, this might fetch data from a database or external service.
	// Here, we return a hardcoded configuration for demonstration purposes.
	return &entity.AppConfig{
		CurrentAppVersion:      "1.0.0",
		MinimumRequiredVersion: "0.9.0",
		LatestAppVersion:       "1.2.0",
		APIBaseURL:             "https://api.example.com",
		Features: entity.Features{
			AIChat:            true,
			BiometricLogin:    true,
			PushNotifications: false,
			DarkMode:          true,
		},
	}, nil
}

func NewAppConfigUsecase() *appConfigUsecase {
	return &appConfigUsecase{}
}
