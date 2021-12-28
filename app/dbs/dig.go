package dbs

import (
	"go.uber.org/dig"
)

// Inject dbs
func Inject(container *dig.Container) error {
	_ = container.Provide(NewDatabase)
	_ = container.Provide(NewRedis)
	return nil
}
