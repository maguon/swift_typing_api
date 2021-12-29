package repos

import (
	"go.uber.org/dig"
)

// Inject repositories
func Inject(container *dig.Container) error {
	_ = container.Provide(NewAppRepo)
	_ = container.Provide(NewAuthRepo)
	_ = container.Provide(NewUserDeviceRepo)
	_ = container.Provide(NewUserRepo)
	_ = container.Provide(NewTPoemRepo)
	_ = container.Provide(NewTWordRepo)
	return nil
}
