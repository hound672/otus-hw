//go:build wireinject
// +build wireinject

package application

import (
	"context"

	"github.com/google/wire"
)

func initApp(ctx context.Context, appConfig *config.AppConfig) (*Application, func(), error) {
	panic(wire.Build(
		wire.Struct(new(Application), "*"),
	))
}
