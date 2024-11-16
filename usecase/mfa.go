package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/ogiogidayo/aconf/config"
	"github.com/ogiogidayo/aconf/utils"
)

type Credentials struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
}

func AuthenticateWithMFA(cfg *config.Config, profile, mfaCode string) {
	mfaArn, exists := cfg.Profiles[profile]
	if !exists {
		fmt.Printf("Profile not found: %s\n", profile)
		return
	}

	output, err := utils.RunShellCommand([]string{
		"aws", "sts", "get-session-token",
		"--serial-number", mfaArn,
		"--profile", profile,
		"--token-code", mfaCode,
	})
	if err != nil {
		fmt.Printf("Error executing AWS CLI: %s\n", err)
		return
	}

	var result struct {
		Credentials Credentials `json:"Credentials"`
	}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		fmt.Printf("Error parsing AWS CLI output: %s\n", err)
		return
	}

	fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", result.Credentials.AccessKeyId)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", result.Credentials.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%s\n", result.Credentials.SessionToken)
}
