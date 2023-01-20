package pac

import "github.com/samber/do"

// Provide will register a service to dependency injector
func Provide[T any](app *App, svc T) {
	do.Provide(app.injector, func(i *do.Injector) (T, error) {
		return svc, nil
	})
}

// Invoke will invoke a service from dependency injector
func Invoke[T any](app *App) T {
	return do.MustInvoke[T](app.injector)
}
