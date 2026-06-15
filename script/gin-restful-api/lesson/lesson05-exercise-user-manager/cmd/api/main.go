package main

import (
	"trungem.com/user-manager/internal/app"
	"trungem.com/user-manager/internal/config"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize application
	application := app.NewApplication(cfg)

	// Start server
	if err := application.Run(); err != nil {
		panic(err)
	}

}
