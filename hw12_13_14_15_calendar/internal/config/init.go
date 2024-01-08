package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
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

	appConfig := &AppConfig{}
	if err := v.Unmarshal(appConfig); err != nil {
		return nil, fmt.Errorf("v.Unmarshal: %w", err)
	}

	return appConfig, nil

	// // logger config
	// viperLogger := v.Sub("logger")
	// loggerConfig := &logger.Config{
	// 	Level: viperLogger.GetString("level"),
	// }
	//
	// // server config
	// viperServer := v.Sub("server")
	// serverConfig := &server.Config{
	// 	UseReflection:         viperServer.GetBool("use_reflection"),
	// 	PortGRPC:              viperServer.GetInt("port_grpc"),
	// 	PortHTTP:              viperServer.GetInt("port_http"),
	// 	HTTPReadTimeout:       viperServer.GetInt("http_read_timeout"),
	// 	HTTPWriteTimeout:      viperServer.GetInt("http_write_timeout"),
	// 	HTTPIdleTimeout:       viperServer.GetInt("http_idle_timeout"),
	// 	HTTPReadHeaderTimeout: viperServer.GetInt("http_read_header_timeout"),
	// }
	//
	// // postgres config
	// viperPostgres := v.Sub("postgres")
	// postgresConfig := &postgres.Config{
	// 	Host:     viperPostgres.GetString("host"),
	// 	Port:     viperPostgres.GetInt("port"),
	// 	Username: viperPostgres.GetString("username"),
	// 	Password: viperPostgres.GetString("password"),
	// 	Database: viperPostgres.GetString("database"),
	// }
	// appConfig := &AppConfig{
	// 	Logger:   loggerConfig,
	// 	Server:   serverConfig,
	// 	Postgres: postgresConfig,
	// }
	// return appConfig, nil
}
