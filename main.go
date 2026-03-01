package main

import (
	"fmt"
	"os"

	"github.com/SShogun/ClawCLI/cmd"
	"github.com/SShogun/ClawCLI/internal/config"
)

func main() {
	// Initialize configuration from .env file
	if err := config.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize config: %v\n", err)
		os.Exit(1)
	}

	cmd.Execute()
}
