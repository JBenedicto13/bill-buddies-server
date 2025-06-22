package container

import (
	"billbuddies/server/app/controller/health"
	"billbuddies/server/app/provider"
	"billbuddies/server/app/repository"
)

// Container holds the dependencies for the application
type Container struct {
	// Providers
	MongoProvider *provider.MongoProvider

	// Repositories
	UserRepository *repository.UserRepository

	// Middleware

	// Controllers
	HealthController *health.HealthController
}

// NewContainer initializes a new IoC container
func NewContainer() *Container {
	// Initialize providers
	mongoProvider := provider.NewMongoProvider()

	// Initialize repositories
	userRepository := repository.NewUserRepository(mongoProvider)

	// Initialize middleware

	// Initialize controllers
	healthController := health.NewHealthController()

	return &Container{
		MongoProvider:    mongoProvider,
		UserRepository:   userRepository,
		HealthController: healthController,
	}
}