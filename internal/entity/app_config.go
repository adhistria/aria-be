package entity

type AppConfig struct {
	CurrentAppVersion      string   `json:"current_app_version"`
	MinimumRequiredVersion string   `json:"minimum_required_version"`
	LatestAppVersion       string   `json:"latest_app_version"`
	APIBaseURL             string   `json:"api_base_url"`
	Features               Features `json:"features"`
}

type Features struct {
	AIChat            bool `json:"ai_chat"`
	BiometricLogin    bool `json:"biometric_login"`
	PushNotifications bool `json:"push_notifications"`
	DarkMode          bool `json:"dark_mode"`
}
