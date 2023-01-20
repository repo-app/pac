package pac

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// ShowFiberBanner will show Fiber banner when app start
func ShowFiberBanner() AppOption {
	return func(a *App) {
		a.BeforeFiberCreated = append(a.BeforeFiberCreated, func(c *fiber.Config) {
			c.DisableStartupMessage = false
		})
	}
}

// ListenPortFromEnv will get listen port from environment variable.
// If environment variable is empty, it will use given default port
func ListenPortFromEnv(default_listen string) AppOption {
	return func(a *App) {
		envPort := os.Getenv("PORT")
		envPort = strings.TrimSpace(envPort)

		a.listenPort = envPort

		if a.listenPort == "" {
			a.listenPort = default_listen
		}
	}
}

// UseLogger will enable Fiber's logger
func UseLogger() AppOption {
	return func(a *App) {
		a.AfterFiberCreated = append(a.AfterFiberCreated, func(a *App) {
			a.fiber.Use(logger.New())
		})
	}
}

// UseCORS will use Fiber's CORS middleware
func UseCORS() AppOption {
	return func(a *App) {
		a.AfterFiberCreated = append(a.AfterFiberCreated, func(a *App) {
			a.fiber.Use(cors.New())
		})
	}
}
