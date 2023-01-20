package pac

import "github.com/gofiber/fiber/v2"

// FiberOption 是 Fiber 的設定選項
type FiberOption func(*fiber.Config)

// AppOption 是應用的設定選項
type AppOption func(app *App)

// Service 是可以被加入到應用中的服務
type Service interface {
	// Register 會在應用啟動時被呼叫
	Register(app *App)
}
