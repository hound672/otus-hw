package config

const configFile = "config.yml"

func Init() (*AppConfig, error) {
	appConfig := &AppConfig{}
	return appConfig, nil
}
