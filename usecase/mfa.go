package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"aconf/config"
	"aconf/utils"
)

type SessionCredentials struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
}

func AuthenticateWithMFA(cfg *config.Config, profile, mfaCode string) error {
	mfaArn, exists := cfg.Profiles[profile]
	if !exists {
		return fmt.Errorf("Profile not found: %s", profile)
	}

	cmd := []string{
		"aws", "sts", "get-session-token",
		"--serial-number", mfaArn,
		"--profile", profile,
		"--token-code", mfaCode,
	}

	output, err := utils.RunShellCommand(cmd)
	if err != nil {
		return err
	}

	var result struct {
		Credentials SessionCredentials `json:"Credentials"`
	}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return err
	}

	os.Setenv("AWS_ACCESS_KEY_ID", result.Credentials.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", result.Credentials.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", result.Credentials.SessionToken)

	fmt.Println("Successfully authenticated with MFA.")
	return nil
}
