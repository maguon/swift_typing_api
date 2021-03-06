package api

import (
	"go.uber.org/dig"
)

// Inject apis
func Inject(container *dig.Container) error {
	_ = container.Provide(NewAppAPI)
	_ = container.Provide(NewUserDeviceAPI)
	_ = container.Provide(NewUserAPI)
	_ = container.Provide(NewTAllAPI)
	return nil
}
