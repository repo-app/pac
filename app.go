// Package pac is a simple web framework based on Fiber, and Do.
// It provides some useful features to make web development easier.
// And dependency injection is the most important feature.
package pac

import (
	"github.com/samber/do"
)

// NewApp return a new Pac application
func NewApp(opts ...AppOption) *App {
	// First, we create a empty application
	pacApp := App{
		services: make([]Service, 0),

		injector:           do.New(),
		BeforeFiberCreated: make([]FiberOption, 0),
		AfterFiberCreated:  make([]AppOption, 0),
	}

	// then, we start to apply all options
	for _, opt := range opts {
		opt(&pacApp)
	}

	// let's create a basic fiber config
	fiberCfg := fiber.Config{
		DisableStartupMessage: true,
	}

	// let's apply all modifies to fiber
	for _, fiberOpt := range pacApp.BeforeFiberCreated {
		fiberOpt(&fiberCfg)
	}

	// create fiber based on config we just prepared
	pacApp.fiber = fiber.New(fiberCfg)

	// then after create fiber, we apply all options in AfterFiberCreated
	for _, opt := range pacApp.AfterFiberCreated {
		opt(&pacApp)
	}

	// after all options applied, we can clear all options
	pacApp.BeforeFiberCreated = nil
	pacApp.AfterFiberCreated = nil

	// return application
	return &pacApp
}

// App is the major runtime program
type App struct {
	// fiber is our HTTP framework
	fiber *fiber.App
	// listenPort is where we listen to
	listenPort string
	// services contains all our services
	services []Service
	// Injector is our dependency injector
	injector *do.Injector

	// hooks
	BeforeFiberCreated []FiberOption
	AfterFiberCreated  []AppOption
}

// Router return fiber's router
func (app *App) Router() *fiber.App {
	return app.fiber
}

// Start will start the application
func (app *App) Start() {
	// before we start, we need to check if listen port is empty
	if app.listenPort == "" {
		panic("[pac] No listen port specified")
	}

	// then start the server
	err := app.fiber.Listen(app.listenPort)

	// if there's error, then panic
	// usually, this is because the port is already in use
	if err != nil {
		panic("[pac] Failed to start the server: " + err.Error())
	}
}

// Add will add service into application, register it to service list
func (app *App) Add(svc Service) {
	svc.Register(app)
}
