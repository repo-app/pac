package pac

import "github.com/samber/do"

// Provide 會將一個服務提供給依賴注入器
func Provide[T any](app *App, svc T) {
	do.Provide(app.injector, func(i *do.Injector) (T, error) {
		return svc, nil
	})
}

// Invoke 會從依賴注入器中取得一個服務
func Invoke[T any](app *App) T {
	return do.MustInvoke[T](app.injector)
}
