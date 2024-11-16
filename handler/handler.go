package cli

import (
	"aconf/config"
	"aconf/usecase"
	"errors"
)

func HandleCommand(command string, args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	switch command {
	case "switch":
		if len(args) != 1 {
			return errors.New("Usage: aconf switch <profile>")
		}
		return usecase.SwitchProfile(cfg, args[0])

	case "add":
		if len(args) != 2 {
			return errors.New("Usage: aconf add <profile> <MFA ARN>")
		}
		return config.AddProfile(cfg, args[0], args[1])

	default:
		if len(args) != 1 {
			return errors.New("Usage: aconf <profile> <MFA code>")
		}
		return usecase.AuthenticateWithMFA(cfg, command, args[0])
	}
}
