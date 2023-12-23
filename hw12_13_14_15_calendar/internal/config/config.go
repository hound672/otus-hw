package config

import (
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/postgres"
)

type AppConfig struct {
	Logger   *logger.Config
	Postgres *postgres.Config
}
