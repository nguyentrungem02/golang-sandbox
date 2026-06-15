package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"trungem.com/shopping-cart/internal/config"
	"trungem.com/shopping-cart/internal/middleware"
	"trungem.com/shopping-cart/internal/routes"
	"trungem.com/shopping-cart/internal/validation"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	if err := validation.InitValidator(); err != nil {
		log.Fatalf("Error initializing validator: %v", err)
	}

	go middleware.CleanupClients()

	loadEnv()

	module := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModuleRoutes(module)...)

	return &Application{
		config:  cfg,
		router:  r,
		modules: module,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}
	return routeList
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
