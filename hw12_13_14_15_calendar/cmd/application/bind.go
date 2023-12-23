package application

import (
	"github.com/google/wire"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/server"
)

// server

var initServer = wire.NewSet(
	wire.FieldsOf(new(*config.AppConfig), "Server"),
	server.New,
)
