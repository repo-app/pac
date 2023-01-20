package pac

import "github.com/gofiber/fiber/v2"

// FiberOption modify Fiber Config
type FiberOption func(*fiber.Config)

// AppOption modify Pac App behavior
type AppOption func(app *App)

// Service can be registered to Pac App
type Service interface {
	// Register will be called when service added into Pac App
	Register(app *App)
}
