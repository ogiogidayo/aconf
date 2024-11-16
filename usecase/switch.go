package usecase

import (
	"aconf/config"
	"os"
)

func SwitchProfile(cfg *config.Config, profile string) error {
	//if _, exists := cfg.Profiles[profile]; !exists {
	//	return errors.New("Profile not found: " + profile)
	//	}

	return os.Setenv("AWS_DEFAULT_PROFILE", profile)
}
