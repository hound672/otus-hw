// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"context"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/config"
)

// Injectors from init.go:

func initApp(ctx context.Context, appConfig *config.AppConfig) (*Application, func(), error) {
	application := &Application{}
	return application, func() {
	}, nil
}
