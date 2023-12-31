package application

import (
	"context"
	"fmt"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/build"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/server"
)

type Application struct {
	server *server.Server
}

func CreateApplication() (*Application, func(), error) {
	ctx := context.Background()
	_ = ctx

	appConfig, err := config.Init()
	if err != nil {
		return nil, nil, err
	}

	if err := logger.InitLogger(appConfig.Logger); err != nil {
		return nil, nil, fmt.Errorf("logger.InitLogger: %w", err)
	}

	app, cleanup, err := initApp(ctx, appConfig)
	if err != nil {
		return nil, nil, err
	}

	stopApplication := func() {
		logger.Info("Stop application")
		cleanup()
		app.server.Stop()
		logger.Info("All done")
	}

	return app, stopApplication, nil
}

func (app *Application) Run() error {
	logger.Info("Start calendar service", "version", build.Version)

	if err := app.server.Run(); err != nil {
		return fmt.Errorf("app.server.Run: %w", err)
	}

	return nil
}
