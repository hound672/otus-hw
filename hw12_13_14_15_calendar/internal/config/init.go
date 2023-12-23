package config

import (
	"os"

	"github.com/spf13/viper"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/postgres"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/server"
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
	viperLogger := v.Sub("logger")
	loggerConfig := &logger.Config{
		Level: viperLogger.GetString("level"),
	}

	// server config
	viperServer := v.Sub("server")
	serverConfig := &server.Config{
		UseReflection: viperServer.GetBool("use_reflection"),
		PortGRPC:      viperServer.GetInt("port_grpc"),
		PortHTTP:      viperServer.GetInt("port_http"),
	}

	// postgres config
	viperPostgres := v.Sub("postgres")
	postgresConfig := &postgres.Config{
		Host:     viperPostgres.GetString("host"),
		Port:     viperPostgres.GetInt("port"),
		Username: viperPostgres.GetString("username"),
		Password: viperPostgres.GetString("password"),
		Database: viperPostgres.GetString("database"),
	}
	appConfig := &AppConfig{
		Logger:   loggerConfig,
		Server:   serverConfig,
		Postgres: postgresConfig,
	}
	return appConfig, nil
}
