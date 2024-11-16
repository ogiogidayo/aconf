package usecase

import (
	"fmt"

	"github.com/ogiogidayo/aconf/config"
)

func SwitchProfile(cfg *config.Config, profile string) error {
	//if _, exists := cfg.Profiles[profile]; !exists {
	//	return errors.New("Profile not found: " + profile)
	//	}

	fmt.Printf("export AWS_DEFAULT_PROFILE=%s\n", profile)
	return nil
}
