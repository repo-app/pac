// 套件 pac 加快我們建立一個服務的基礎，包含了依賴注入器、路由及一些常用的方法
package pac

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

// NewApp 回傳一個新的 Pac 應用
func NewApp(opts ...AppOption) *App {
	pacApp := App{
		services: make([]Service, 0),

		injector:           do.New(),
		BeforeFiberCreated: make([]FiberOption, 0),
		AfterFiberCreated:  make([]AppOption, 0),
	}

	// 執行所有的 AppOption
	for _, opt := range opts {
		opt(&pacApp)
	}

	// 建立 Fiber 設定
	fiberCfg := fiber.Config{
		DisableStartupMessage: true,
	}

	// 執行所有的 FiberOption 設定
	for _, fiberOpt := range pacApp.BeforeFiberCreated {
		fiberOpt(&fiberCfg)
	}

	// 建立 fiber
	pacApp.fiber = fiber.New(fiberCfg)

	// 執行所有的 AppOption Hook 設定
	for _, opt := range pacApp.AfterFiberCreated {
		opt(&pacApp)
	}

	// 清空所有的 hook
	pacApp.BeforeFiberCreated = nil
	pacApp.AfterFiberCreated = nil

	// 回傳組合好的應用
	return &pacApp
}

// App 是整個執行時應用的主要結構
type App struct {
	// fiber 是我們的 web 框架
	fiber *fiber.App
	// listenPort 是我們要監聽的位置，通常是 :3000
	listenPort string
	// services 是我們的服務列表
	services []Service
	// Injector 是我們的依賴注入器
	injector *do.Injector

	// 一些在應用啟動時會被呼叫的函式
	BeforeFiberCreated []FiberOption
	AfterFiberCreated  []AppOption
}

// Router 回傳程式內部的路由器
func (app *App) Router() *fiber.App {
	return app.fiber
}

// Start 啟動應用並開始監聽
func (app *App) Start() {
	// 檢查我們是否有監聽位置
	if app.listenPort == "" {
		panic("[pac] No listen port specified")
	}

	// 開始監聽
	err := app.fiber.Listen(app.listenPort)

	// 如果有錯誤，則停止應用
	if err != nil {
		panic("[pac] Failed to start the server: " + err.Error())
	}
}

// Add 會將服務加入到應用中
func (app *App) Add(svc Service) {
	svc.Register(app)
}
