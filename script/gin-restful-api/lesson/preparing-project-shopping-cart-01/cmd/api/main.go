package main

import (
	"trungem.com/shopping-cart/internal/app"
	"trungem.com/shopping-cart/internal/config"
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
