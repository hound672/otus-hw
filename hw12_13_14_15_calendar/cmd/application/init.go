//go:build wireinject
// +build wireinject

package application

import (
	"context"

	"github.com/google/wire"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/config"
)

func initApp(ctx context.Context, appConfig *config.AppConfig) (*Application, func(), error) {
	panic(wire.Build(
		wire.Struct(new(Application), "*"),
	))
}
