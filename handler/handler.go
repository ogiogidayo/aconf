package handler

import (
	"fmt"
	"os"

	"github.com/ogiogidayo/aconf/config"
	"github.com/ogiogidayo/aconf/usecase"
)

func HandleCommand(command string, args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	switch command {
	case "switch":
		if len(args) != 1 {
			fmt.Println("Usage: aconf switch <profile>")
			os.Exit(1)
		}
		if err := usecase.SwitchProfile(cfg, args[0]); err != nil {
			return fmt.Errorf("error in usecase: %v", err)
		}

	case "add":
		if len(args) != 2 {
			fmt.Println("Usage: aconf add <profile> <MFA ARN>")
			os.Exit(1)
		}
		if err := config.AddProfile(cfg, args[0], args[1]); err != nil {
			return fmt.Errorf("error in usecase: %v", err)
		}

	default:
		if len(args) != 1 {
			fmt.Println("Usage: aconf <profile> <MFA code>")
			os.Exit(1)
		}
		usecase.AuthenticateWithMFA(cfg, command, args[0])
	}
	return nil
}
