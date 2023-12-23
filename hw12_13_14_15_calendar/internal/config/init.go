package config

import (
	"os"

	"github.com/spf13/viper"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/postgres"
)

const configFile = "config.yml"

func Init() (*AppConfig, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		if value == "" {
			continue
		}
		v.Set(k, os.ExpandEnv(value))
	}

	// logger config
	viperLogger := viper.Sub("logger")
	loggerConfig := &logger.Config{
		Level: viperLogger.GetString("level"),
	}

	// postgres config
	viperPostgres := viper.Sub("postgres")
	postgresConfig := &postgres.Config{
		Host:     viperPostgres.GetString("host"),
		Port:     viperPostgres.GetInt("port"),
		Username: viperPostgres.GetString("username"),
		Password: viperPostgres.GetString("password"),
		Database: viperPostgres.GetString("database"),
	}
	appConfig := &AppConfig{
		Logger:   loggerConfig,
		Postgres: postgresConfig,
	}
	return appConfig, nil
}
