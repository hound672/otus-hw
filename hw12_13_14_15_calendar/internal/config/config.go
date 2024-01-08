package config

import (
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/postgres"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/server"
)

type AppConfig struct {
	Logger   *logger.Config
	Server   *server.Config
	Postgres *postgres.Config
}
