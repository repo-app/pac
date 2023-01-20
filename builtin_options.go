package pac

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// ShowFiberBanner 會使得應用在啟動時顯示 Fiber 的啟動訊息
func ShowFiberBanner() func(a *App) {
	return func(a *App) {
		a.BeforeFiberCreated = append(a.BeforeFiberCreated, func(c *fiber.Config) {
			c.DisableStartupMessage = false
		})
	}
}

// ListenPortFromEnv 會從環境變數中取得監聽的位置，如果沒有則用預設值
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

// UseLogger 會在應用啟動時啟用 Fiber 的 Logger
func UseLogger() AppOption {
	return func(a *App) {
		a.AfterFiberCreated = append(a.AfterFiberCreated, func(a *App) {
			a.fiber.Use(logger.New())
		})
	}
}

// UseCORS 會在應用啟動時啟用 Fiber 的 CORS
func UseCORS() AppOption {
	return func(a *App) {
		a.AfterFiberCreated = append(a.AfterFiberCreated, func(a *App) {
			a.fiber.Use(cors.New())
		})
	}
}
