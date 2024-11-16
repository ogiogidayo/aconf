package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aconf <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	if err := handler.HandleCommand(command, args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
